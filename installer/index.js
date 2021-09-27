const mkdirp = require("mkdirp");
const zlib = require("zlib");
const tar = require("tar");

const VERSION_NUMBER = "v0.0.1";

const PLATFORM_MAPPING = {
    "darwin" : "darwin",
    "linux" : "linux",
    "win32" : "windows",
    "freebsd" : "freebsd"
}

const ARCH_MAPPING = {
    "ia32": "386",
    "x64": "amd64",
    "arm": "arm"
}

async function getInstallationPath() {
    await exec("npm bin", function(err, stdout, stderr) {
        let dir = null
        if (err || stderr || !stdout || stdout.length === 0) {
            let env = process.env;
            if (env && env.npm_config_prefix) {
                dir = path.join(env.npm_config_prefix, "bin")
            }
        }else{
            dir = stdout.trim()
        }

        mkdirp.sync(dir)
        return dir;
    })
}

async function installer() {
    const arch = process.arch
    const platform = process.platform
    
    if (!(arch in ARCH_MAPPING)) {
        console.error("Installation not supported for the architecture: " + arch);
        return;
    }
    if (!(platform in PLATFORM_MAPPING)) {
        console.error("Installation not supported for this platform: " + platform);
        return;
    }

    let binName = "bare";
    let binPath = "./barego/bin";
    let url = `https://github.com/madrix01/bare/releases/tag/${VERSION_NUMBER}`

    mkdirp.sync(binPath)

    let ungz = zlib.createGunzip();
    let untar = tar.Extract({path : binPath});

    
}

installer()