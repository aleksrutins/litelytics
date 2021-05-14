export function resolveRel(filename, meta) {
    return new URL(filename, meta.url).pathname;
}