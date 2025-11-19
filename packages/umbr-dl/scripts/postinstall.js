const fs = require("fs");
const path = require("path");

const binDir = path.join(__dirname, "..", "bin");

const binaries = [
  "go-download-linux",
  "go-download-darwin",
  "go-download-windows.exe",
];

binaries.forEach((bin) => {
  const filePath = path.join(binDir, bin);

  if (fs.existsSync(filePath)) {
    try {
      // Windows doesn't need chmod
      if (process.platform !== "win32") {
        fs.chmodSync(filePath, 0o755);
        console.log(`Set executable permissions: ${bin}`);
      } else {
        console.log(`Skipping chmod on Windows: ${bin}`);
      }
    } catch (err) {
      console.error(`Failed to set permissions on ${bin}:`, err);
    }
  }
});
