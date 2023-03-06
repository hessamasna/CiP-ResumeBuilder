// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
    app: {
        head: {
            htmlAttrs: {dir: 'rtl', lang: 'fa'},
        },
    },
    modules: ['@nuxtjs/tailwindcss'],
    css: ['~/assets/css/main.css', 'vuetify/lib/styles/main.sass', '@mdi/font/css/materialdesignicons.min.css'],
    build: {
        transpile: ['vuetify','@headlessui/vue'],
    },
})
