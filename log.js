export function log(msg) {
    if(!log.nocolor) console.log("\x1B[34m➤\x1B[0m " + msg);
    else console.log("➤ " + msg);
}
log.nocolor = false;