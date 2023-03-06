import {createStore} from "vuex";

const store = createStore({
    state() {
        return {
            status: {
                isLoggedIn: false,
                access_token: '',
                refresh_token: '',
                Email: ''
            },
            loginData: {},
        };
    },
    mutations: {
        setStatus(state, data) {
            console.log(data)
            state.status = data;
        },
        setLoginData(state, data) {
            state.loginData = data;
        }
    },
});

export default defineNuxtPlugin((nuxtApp) => {
    nuxtApp.vueApp.use(store);
    // Install the store instance as a plugin

});
