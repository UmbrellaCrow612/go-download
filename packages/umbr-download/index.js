const { spawn } = require("child_process");
const path = require("path");
const fs = require("fs");

/**
 * Resolves the correct Go binary for the current OS.
 *
 * The expected binary names are:
 * - `go-download-linux`
 * - `go-download-darwin`
 * - `go-download-windows.exe`
 *
 * The binaries must exist in the `bin/` directory next to this file.
 *
 * @returns {string} Absolute path to the platform-specific binary.
 * @throws {Error} If the platform is unsupported or binary not found.
 */
function getPlatformBinary() {
  const platform = process.platform;

  let binaryName;

  if (platform === "linux") {
    binaryName = "go-download-linux";
  } else if (platform === "darwin") {
    binaryName = "go-download-darwin";
  } else if (platform === "win32") {
    binaryName = "go-download-windows.exe";
  } else {
    throw new Error(`Unsupported platform: ${platform}`);
  }

  const binaryPath = path.join(__dirname, "bin", binaryName);

  if (!fs.existsSync(binaryPath)) {
    throw new Error(`Binary not found: ${binaryPath}`);
  }

  return binaryPath;
}

/**
 * Downloads a file using the platform-specific Go binary.
 *
 * @typedef {Object} DownloadOptions
 * @property {string} path - Directory to save the file.
 * @property {string} [name] - Optional new filename.
 * @property {number} [timeout] - Optional timeout in milliseconds. If exceeded, the download is aborted.
 *
 * @param {string} url - The URL to download from.
 * @param {DownloadOptions} opts - Download configuration.
 *
 * @returns {Promise<void>} Resolves on success, rejects on failure or timeout.
 */
function download(url, opts = {
    timeout: 6000,
    path: ""
}) {
  if (!url) {
    return Promise.reject(new Error("url is required"));
  }
  if (!opts.path) {
    return Promise.reject(new Error("path is required in options"));
  }

  const binary = getPlatformBinary();

  const args = [url, opts.path];

  if (opts.name) {
    args.push("--name", opts.name);
  }

  return new Promise((resolve, reject) => {
    const child = spawn(binary, args, {
      stdio: "inherit",
    });

    let timeoutId = null;

    if (typeof opts.timeout === "number" && opts.timeout > 0) {
      timeoutId = setTimeout(() => {
        child.kill("SIGKILL");
        reject(new Error(`Download timed out after ${opts.timeout} ms`));
      }, opts.timeout);
    }

    child.on("error", (err) => {
      if (timeoutId) clearTimeout(timeoutId);
      reject(err);
    });

    child.on("close", (code) => {
      if (timeoutId) clearTimeout(timeoutId);

      if (code === 0) {
        resolve();
      } else {
        reject(new Error(`Download process exited with code ${code}`));
      }
    });
  });
}

module.exports = {
  download,
};
