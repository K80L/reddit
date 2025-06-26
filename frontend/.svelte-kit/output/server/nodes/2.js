

export const index = 2;
let component_cache;
export const component = async () => component_cache ??= (await import('../entries/pages/_page.svelte.js')).default;
export const imports = ["_app/immutable/nodes/2.YOfeqm2G.js","_app/immutable/chunks/scheduler.Bmg8oFKD.js","_app/immutable/chunks/index.CBPgh3uM.js"];
export const stylesheets = ["_app/immutable/assets/2.D1U7sR_4.css"];
export const fonts = [];
