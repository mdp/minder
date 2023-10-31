"use strict";(self.webpackChunkstacklok=self.webpackChunkstacklok||[]).push([[5270],{3905:(e,r,t)=>{t.d(r,{Zo:()=>s,kt:()=>v});var n=t(7294);function o(e,r,t){return r in e?Object.defineProperty(e,r,{value:t,enumerable:!0,configurable:!0,writable:!0}):e[r]=t,e}function i(e,r){var t=Object.keys(e);if(Object.getOwnPropertySymbols){var n=Object.getOwnPropertySymbols(e);r&&(n=n.filter((function(r){return Object.getOwnPropertyDescriptor(e,r).enumerable}))),t.push.apply(t,n)}return t}function a(e){for(var r=1;r<arguments.length;r++){var t=null!=arguments[r]?arguments[r]:{};r%2?i(Object(t),!0).forEach((function(r){o(e,r,t[r])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(t)):i(Object(t)).forEach((function(r){Object.defineProperty(e,r,Object.getOwnPropertyDescriptor(t,r))}))}return e}function l(e,r){if(null==e)return{};var t,n,o=function(e,r){if(null==e)return{};var t,n,o={},i=Object.keys(e);for(n=0;n<i.length;n++)t=i[n],r.indexOf(t)>=0||(o[t]=e[t]);return o}(e,r);if(Object.getOwnPropertySymbols){var i=Object.getOwnPropertySymbols(e);for(n=0;n<i.length;n++)t=i[n],r.indexOf(t)>=0||Object.prototype.propertyIsEnumerable.call(e,t)&&(o[t]=e[t])}return o}var p=n.createContext({}),c=function(e){var r=n.useContext(p),t=r;return e&&(t="function"==typeof e?e(r):a(a({},r),e)),t},s=function(e){var r=c(e.components);return n.createElement(p.Provider,{value:r},e.children)},d="mdxType",u={inlineCode:"code",wrapper:function(e){var r=e.children;return n.createElement(n.Fragment,{},r)}},m=n.forwardRef((function(e,r){var t=e.components,o=e.mdxType,i=e.originalType,p=e.parentName,s=l(e,["components","mdxType","originalType","parentName"]),d=c(t),m=o,v=d["".concat(p,".").concat(m)]||d[m]||u[m]||i;return t?n.createElement(v,a(a({ref:r},s),{},{components:t})):n.createElement(v,a({ref:r},s))}));function v(e,r){var t=arguments,o=r&&r.mdxType;if("string"==typeof e||o){var i=t.length,a=new Array(i);a[0]=m;var l={};for(var p in r)hasOwnProperty.call(r,p)&&(l[p]=r[p]);l.originalType=e,l[d]="string"==typeof e?e:o,a[1]=l;for(var c=2;c<i;c++)a[c]=t[c];return n.createElement.apply(null,a)}return n.createElement.apply(null,t)}m.displayName="MDXCreateElement"},3005:(e,r,t)=>{t.r(r),t.d(r,{assets:()=>p,contentTitle:()=>a,default:()=>u,frontMatter:()=>i,metadata:()=>l,toc:()=>c});var n=t(7462),o=(t(7294),t(3905));const i={title:"minder auth revoke provider"},a=void 0,l={unversionedId:"ref/cli/minder_auth_revoke_provider",id:"ref/cli/minder_auth_revoke_provider",title:"minder auth revoke provider",description:"minder auth revoke_provider",source:"@site/docs/ref/cli/minder_auth_revoke_provider.md",sourceDirName:"ref/cli",slug:"/ref/cli/minder_auth_revoke_provider",permalink:"/ref/cli/minder_auth_revoke_provider",draft:!1,tags:[],version:"current",frontMatter:{title:"minder auth revoke provider"},sidebar:"mediator",previous:{title:"minder auth logout",permalink:"/ref/cli/minder_auth_logout"},next:{title:"minder completion",permalink:"/ref/cli/minder_completion"}},p={},c=[{value:"minder auth revoke_provider",id:"minder-auth-revoke_provider",level:2},{value:"Synopsis",id:"synopsis",level:3},{value:"Options",id:"options",level:3},{value:"Options inherited from parent commands",id:"options-inherited-from-parent-commands",level:3},{value:"SEE ALSO",id:"see-also",level:3}],s={toc:c},d="wrapper";function u(e){let{components:r,...t}=e;return(0,o.kt)(d,(0,n.Z)({},s,t,{components:r,mdxType:"MDXLayout"}),(0,o.kt)("h2",{id:"minder-auth-revoke_provider"},"minder auth revoke_provider"),(0,o.kt)("p",null,"Revoke access tokens for provider"),(0,o.kt)("h3",{id:"synopsis"},"Synopsis"),(0,o.kt)("p",null,"It can revoke access tokens for specific provider."),(0,o.kt)("pre",null,(0,o.kt)("code",{parentName:"pre"},"minder auth revoke_provider [flags]\n")),(0,o.kt)("h3",{id:"options"},"Options"),(0,o.kt)("pre",null,(0,o.kt)("code",{parentName:"pre"},"  -a, --all                 Revoke all tokens\n  -h, --help                help for revoke_provider\n  -g, --project-id string   ID of the project for repo registration\n  -n, --provider string     Name for the provider to revoke tokens for\n")),(0,o.kt)("h3",{id:"options-inherited-from-parent-commands"},"Options inherited from parent commands"),(0,o.kt)("pre",null,(0,o.kt)("code",{parentName:"pre"},'      --config string            Config file (default is $PWD/config.yaml)\n      --grpc-host string         Server host (default "staging.stacklok.dev")\n      --grpc-insecure            Allow establishing insecure connections\n      --grpc-port int            Server port (default 443)\n      --identity-client string   Identity server client ID (default "mediator-cli")\n      --identity-realm string    Identity server realm (default "stacklok")\n      --identity-url string      Identity server issuer URL (default "https://auth.staging.stacklok.dev")\n')),(0,o.kt)("h3",{id:"see-also"},"SEE ALSO"),(0,o.kt)("ul",null,(0,o.kt)("li",{parentName:"ul"},(0,o.kt)("a",{parentName:"li",href:"/ref/cli/minder_auth"},"minder auth"),"\t - Authorize and manage accounts within a minder control plane")))}u.isMDXComponent=!0}}]);