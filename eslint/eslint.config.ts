import { withNuxt } from "./.nuxt/eslint.config.mjs";

export default withNuxt({
  rules: {
    //   79:13  warning  Disallow self-closing on HTML void elements (<input/>)  vue/html-self-closing
    "vue/html-self-closing": [
      "off",
      {
        html: {
          void: "any", // allow self-closing on void elements
          normal: "never",
          component: "always",
        },
        svg: "always",
        math: "always",
      },
    ],
  },
});
