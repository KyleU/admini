(()=>{function m(){for(let e of Array.from(document.querySelectorAll(".menu-container .final")))e.scrollIntoView({block:"nearest"})}function f(){let e=document.getElementById("flash-container");if(e===null)return;let n=e.querySelectorAll(".flash");n.length>0&&setTimeout(()=>{for(let o of n){let r=o;r.style.opacity="0",setTimeout(()=>r.remove(),500)}},3e3)}var u="--selected";function d(){let e={},n={};for(let o of Array.from(document.querySelectorAll(".editor"))){let r=o,s=()=>{e={},n={};for(let l of r.elements){let t=l;if(t.name.length>0)if(t.name.endsWith(u))n[t.name]=t;else{(t.type!=="radio"||t.checked)&&(e[t.name]=t.value);let a=()=>{let i=n[t.name+u];i&&(i.checked=e[t.name]!==t.value)};t.onchange=a,t.onkeyup=a}}};r.onreset=s,s()}}function p(){if(window.Sortable)for(let e of Array.from(document.getElementsByClassName("sortable")))y(e)}function E(e){for(;e.parentElement&&!e.classList.contains("drag-container");)e=e.parentElement;e.classList.remove("readonly"),y(e)}function y(e){let n=window.Sortable;if(n){let o=e.querySelector(".l");o||(o=e);let s={group:{name:"nested"},handle:".handle",onAdd:t=>{let a=t.item;new n(a.querySelector(".container"),s),a.querySelector(".remove").onclick=function(){g(e,a)},c(e)},onUpdate:()=>c(e),animation:150,fallbackOnBody:!0,swapThreshold:.65};for(let t of Array.from(o.getElementsByClassName("container")))new n(t,s);for(let t of Array.from(o.getElementsByClassName("remove")))t.onclick=function(){g(e,t.parentElement?.parentElement)};let l=e.querySelector(".r");if(l){let t={group:{name:"nested",pull:"clone",put:!1},handle:".handle",animation:150,fallbackOnBody:!0,swapThreshold:.65,sort:!1};for(let a of Array.from(l.getElementsByClassName("container")))new n(a,t)}c(e)}}function g(e,n){n.remove(),c(e)}function c(e){let n=document.querySelector(".drag-state");if(!n)return;let o=document.querySelector(".drag-state-original"),r=e.querySelector(".tracked"),[s,l]=h(r),t=JSON.stringify(s);if(o){o.value.length===0&&(o.value=t);let a=document.querySelector(".drag-actions");a&&(o.value===t?a.classList.add("no-changes"):a.classList.remove("no-changes"));let i=document.querySelector(".drag-tracked-size");i&&(i.innerText=l.toString(10))}n.value=t}function h(e){let n=0,o=[];for(let r of Array.from(e.children))if(r.classList.contains("item")){let[s,l]=L(r);s&&o.push(s),n+=l}return[o,n]}function L(e){let n=1,o={k:e.dataset.key,p:e.dataset.originalPath};for(let r of Array.from(e.children))if(r.classList.contains("container")){let[s,l]=h(r);s.length>0&&(o.c=s),n+=l}return[o,n]}function I(){for(let e of Array.from(document.getElementsByClassName("link-confirm"))){let n=e;n.onclick=function(){let o=n.dataset.message;return o&&o.length===0&&(o="Are you sure?"),confirm(o)}}}function k(){window.admini={sortableEdit:E},m(),f(),I(),d(),p()}document.addEventListener("DOMContentLoaded",k);})();
//# sourceMappingURL=client.js.map
