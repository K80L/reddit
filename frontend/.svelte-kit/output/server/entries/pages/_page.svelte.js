import { c as create_ssr_component, v as validate_component } from "../../chunks/ssr.js";
const css$1 = {
  code: ".card.svelte-1enc9gg{border-bottom:1px solid #27282a;min-height:100px;width:100%;padding:0.25rem 1rem}",
  map: '{"version":3,"file":"card.svelte","sources":["card.svelte"],"sourcesContent":["<div class=\\"card\\"></div>\\n\\n<style>\\n  .card {\\n    border-bottom: 1px solid #27282a;\\n    min-height: 100px;\\n    width: 100%;\\n    padding: 0.25rem 1rem;\\n  }\\n</style>\\n"],"names":[],"mappings":"AAGE,oBAAM,CACJ,aAAa,CAAE,GAAG,CAAC,KAAK,CAAC,OAAO,CAChC,UAAU,CAAE,KAAK,CACjB,KAAK,CAAE,IAAI,CACX,OAAO,CAAE,OAAO,CAAC,IACnB"}'
};
const Card = create_ssr_component(($$result, $$props, $$bindings, slots) => {
  $$result.css.add(css$1);
  return `<div class="card svelte-1enc9gg"></div>`;
});
const css = {
  code: ".main-content.svelte-17djp40{grid-column:3 / span 8;display:flex;flex-direction:column;align-items:center}",
  map: '{"version":3,"file":"+page.svelte","sources":["+page.svelte"],"sourcesContent":["<script lang=\\"ts\\">import { onMount } from \\"svelte\\";\\nimport Card from \\"../components/card.svelte\\";\\nonMount(async () => {\\n  try {\\n    const response = await fetch(\\"http://localhost:8080/api/post\\", {});\\n    const data = await response.json();\\n    console.log(data);\\n  } catch (error) {\\n    console.error(error);\\n  }\\n});\\n<\/script>\\n\\n<main class=\\"main-content\\">\\n  <Card />\\n  <div>Hello</div>\\n</main>\\n\\n<style>\\n  .main-content {\\n    grid-column: 3 / span 8;\\n    display: flex;\\n    flex-direction: column;\\n    align-items: center;\\n  }\\n</style>\\n"],"names":[],"mappings":"AAmBE,4BAAc,CACZ,WAAW,CAAE,CAAC,CAAC,CAAC,CAAC,IAAI,CAAC,CAAC,CACvB,OAAO,CAAE,IAAI,CACb,cAAc,CAAE,MAAM,CACtB,WAAW,CAAE,MACf"}'
};
const Page = create_ssr_component(($$result, $$props, $$bindings, slots) => {
  $$result.css.add(css);
  return `<main class="main-content svelte-17djp40">${validate_component(Card, "Card").$$render($$result, {}, {}, {})} <div data-svelte-h="svelte-k7y22y">Hello</div> </main>`;
});
export {
  Page as default
};
