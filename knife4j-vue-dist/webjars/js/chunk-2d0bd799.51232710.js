(window.webpackJsonp=window.webpackJsonp||[]).push([["chunk-2d0bd799"],{"2bc6":function(t,e,a){"use strict";a.r(e);var n=(a("14d9"),{name:"DataType",props:{text:{type:String,default:"string",required:!0},record:{type:Object,required:!0}},data:function(){return{validators:[]}},created:function(){this.intiValidator()},methods:{intiValidator:function(){var t=this.record;if(null!=t.validateInstance)for(var e in this.getJsonKeyLength(t.validateInstance),t.validateInstance){var a=e+":"+t.validateInstance[e];this.validators.push({key:e,val:a})}},getJsonKeyLength:function(t){var e=0;if(null!=t)for(var a in t)t.hasOwnProperty(a)&&e++;return e}}}),i=a("1805"),r=Object(i.a)(n,(function(){var t=this,e=t._self._c;return e("div",[t.record.validateStatus?e("span",{staticClass:"knife4j-request-validate-jsr"},[e("a-tooltip",{attrs:{placement:"right"}},[e("template",{slot:"title"},t._l(t.validators,(function(a){return e("div",{key:a.key},[t._v(t._s(a.val))])})),0),t._v(" "+t._s(t.text)+" ")],2)],1):e("span",[t._v(t._s(null==t.text||""==t.text?"string":t.text))])])}),[],!1,null,null,null);e.default=r.exports}}]);