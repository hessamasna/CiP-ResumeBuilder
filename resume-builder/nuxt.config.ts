// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
    app: {
        head: {
            htmlAttrs: {dir: 'rtl', lang: 'fa'},
        },
    },
    modules: ['@nuxtjs/tailwindcss', '@element-plus/nuxt'],
    css: ['~/assets/css/main.css'],
})
