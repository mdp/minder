"use strict";(self.webpackChunkstacklok=self.webpackChunkstacklok||[]).push([[7406],{3905:(e,t,r)=>{r.d(t,{Zo:()=>s,kt:()=>m});var n=r(7294);function i(e,t,r){return t in e?Object.defineProperty(e,t,{value:r,enumerable:!0,configurable:!0,writable:!0}):e[t]=r,e}function o(e,t){var r=Object.keys(e);if(Object.getOwnPropertySymbols){var n=Object.getOwnPropertySymbols(e);t&&(n=n.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),r.push.apply(r,n)}return r}function l(e){for(var t=1;t<arguments.length;t++){var r=null!=arguments[t]?arguments[t]:{};t%2?o(Object(r),!0).forEach((function(t){i(e,t,r[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(r)):o(Object(r)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(r,t))}))}return e}function p(e,t){if(null==e)return{};var r,n,i=function(e,t){if(null==e)return{};var r,n,i={},o=Object.keys(e);for(n=0;n<o.length;n++)r=o[n],t.indexOf(r)>=0||(i[r]=e[r]);return i}(e,t);if(Object.getOwnPropertySymbols){var o=Object.getOwnPropertySymbols(e);for(n=0;n<o.length;n++)r=o[n],t.indexOf(r)>=0||Object.prototype.propertyIsEnumerable.call(e,r)&&(i[r]=e[r])}return i}var c=n.createContext({}),a=function(e){var t=n.useContext(c),r=t;return e&&(r="function"==typeof e?e(t):l(l({},t),e)),r},s=function(e){var t=a(e.components);return n.createElement(c.Provider,{value:t},e.children)},y="mdxType",d={inlineCode:"code",wrapper:function(e){var t=e.children;return n.createElement(n.Fragment,{},t)}},u=n.forwardRef((function(e,t){var r=e.components,i=e.mdxType,o=e.originalType,c=e.parentName,s=p(e,["components","mdxType","originalType","parentName"]),y=a(r),u=i,m=y["".concat(c,".").concat(u)]||y[u]||d[u]||o;return r?n.createElement(m,l(l({ref:t},s),{},{components:r})):n.createElement(m,l({ref:t},s))}));function m(e,t){var r=arguments,i=t&&t.mdxType;if("string"==typeof e||i){var o=r.length,l=new Array(o);l[0]=u;var p={};for(var c in t)hasOwnProperty.call(t,c)&&(p[c]=t[c]);p.originalType=e,p[y]="string"==typeof e?e:i,l[1]=p;for(var a=2;a<o;a++)l[a]=r[a];return n.createElement.apply(null,l)}return n.createElement.apply(null,r)}u.displayName="MDXCreateElement"},605:(e,t,r)=>{r.r(t),r.d(t,{assets:()=>c,contentTitle:()=>l,default:()=>d,frontMatter:()=>o,metadata:()=>p,toc:()=>a});var n=r(7462),i=(r(7294),r(3905));const o={},l=void 0,p={unversionedId:"cli/medic_policy_type_list",id:"cli/medic_policy_type_list",title:"medic_policy_type_list",description:"medic policy_type list",source:"@site/docs/cli/medic_policy_type_list.md",sourceDirName:"cli",slug:"/cli/medic_policy_type_list",permalink:"/cli/medic_policy_type_list",draft:!1,tags:[],version:"current",frontMatter:{},sidebar:"mediator",previous:{title:"medic_policy_type_get",permalink:"/cli/medic_policy_type_get"},next:{title:"medic_policy_violation",permalink:"/cli/medic_policy_violation"}},c={},a=[{value:"medic policy_type list",id:"medic-policy_type-list",level:2},{value:"Synopsis",id:"synopsis",level:3},{value:"Options",id:"options",level:3},{value:"Options inherited from parent commands",id:"options-inherited-from-parent-commands",level:3},{value:"SEE ALSO",id:"see-also",level:3},{value:"Auto generated by spf13/cobra on 14-Jul-2023",id:"auto-generated-by-spf13cobra-on-14-jul-2023",level:6}],s={toc:a},y="wrapper";function d(e){let{components:t,...r}=e;return(0,i.kt)(y,(0,n.Z)({},s,r,{components:t,mdxType:"MDXLayout"}),(0,i.kt)("h2",{id:"medic-policy_type-list"},"medic policy_type list"),(0,i.kt)("p",null,"List policy types for a provider within a mediator control plane"),(0,i.kt)("h3",{id:"synopsis"},"Synopsis"),(0,i.kt)("p",null,"The medic policy_type list subcommand lets you list policy types within a\nmediator control plane for an specific provider."),(0,i.kt)("pre",null,(0,i.kt)("code",{parentName:"pre"},"medic policy_type list [flags]\n")),(0,i.kt)("h3",{id:"options"},"Options"),(0,i.kt)("pre",null,(0,i.kt)("code",{parentName:"pre"},"  -h, --help              help for list\n  -o, --output string     Output format (json or yaml)\n  -p, --provider string   Provider to list policies for\n")),(0,i.kt)("h3",{id:"options-inherited-from-parent-commands"},"Options inherited from parent commands"),(0,i.kt)("pre",null,(0,i.kt)("code",{parentName:"pre"},'      --config string      config file (default is $PWD/config.yaml)\n      --grpc-host string   Server host (default "localhost")\n      --grpc-port int      Server port (default 8090)\n')),(0,i.kt)("h3",{id:"see-also"},"SEE ALSO"),(0,i.kt)("ul",null,(0,i.kt)("li",{parentName:"ul"},(0,i.kt)("a",{parentName:"li",href:"/cli/medic_policy_type"},"medic policy_type"),"\t - Manage policy types within a mediator control plane")),(0,i.kt)("h6",{id:"auto-generated-by-spf13cobra-on-14-jul-2023"},"Auto generated by spf13/cobra on 14-Jul-2023"))}d.isMDXComponent=!0}}]);