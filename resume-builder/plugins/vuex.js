import {createStore} from "vuex";

const store = createStore({
    state() {
        return {
            status: {
                isLoggedIn: false,
                access_token: '',
                refresh_token: '',
                Email: '',
                id: 2,
            },
            cookie: '',
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
        },
        setCookie(state, data) {
            state.cookie = data;
        }
    },
});

export default defineNuxtPlugin((nuxtApp) => {
    nuxtApp.vueApp.use(store);
    // Install the store instance as a plugin
});
