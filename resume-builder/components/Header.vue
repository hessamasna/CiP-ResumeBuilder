<template>
  <Disclosure as="nav" class="" v-slot="{ open }">
    <div class="w-full shadow">

      <div class="mx-auto max-w-7xl px-2 sm:px-6 lg:px-8">
        <div class="relative flex h-16 items-center justify-between center">
          <div class="absolute inset-y-0 left-0 flex items-center sm:hidden">
            <!-- Mobile menu button-->
<!--            <DisclosureButton-->
<!--                class="inline-flex items-center justify-center rounded-md p-2 text-gray-400 hover:bg-gray-700 hover:text-white focus:outline-none focus:ring-2 focus:ring-inset focus:ring-white">-->
<!--              <span class="sr-only">Open main menu</span>-->
<!--              <Bars3Icon v-if="!open" class="block h-6 w-6" aria-hidden="true"/>-->
<!--              <XMarkIcon v-else class="block h-6 w-6" aria-hidden="true"/>-->
<!--            </DisclosureButton>-->
          </div>
          <div class="flex flex-1 items-center sm:gap-28 sm:justify-start ">
            <div class="flex flex-shrink-0 items-center">
              <NuxtLink to="/">
                <img src="/img/logo.png" class="h-16"/>
              </NuxtLink>
            </div>
            <div class="hidden sm:ml-6 sm:block  ">
              <div class="flex items-center gap-4 ">
                <NuxtLink v-for="item in navigation" :key="item.name" :to="item.href"
                          @click="toggleNavigationActivation(item)"
                          :class="[item.current ? 'text-green-500 border-b-2 border-green-500 ' : 'hover:text-green-500 hover:border-b-2', 'px-3 py-5  text-sm font-medium border-green-500 text-gray-400']"
                >{{ item.name }}
                </NuxtLink>
              </div>
            </div>
          </div>
          <div class="flex flex-row">
            <Menu as="div" class="relative ml-3" @click="toggleLoginPopUp()" v-if="!isLogin">
              <div class="flex gap-2">

                <div
                    class="text-green-500 border-green-500 border px-8 py-1 rounded-xl hover:cursor-pointer ">
                  ورود
                </div>

              </div>
            </Menu>
            <Menu as="div" class="relative ml-3" @click="toggleSignupPopUp()" v-if="!isLogin">
              <div>
                <div
                    class="text-white bg-green-500 border-green-500 border px-8 py-1 rounded-xl hover:cursor-pointer ">
                  ثبت نام
                </div>
              </div>
            </Menu>

            <div v-if="isLogin" class="text-green-500 font-semibold flex">
              <div
                  class="text-green-500 px-4 py-1  hover:underline hover:cursor-pointer ">
                <NuxtLink to="/dashboard">
                  رزومه های من
                </NuxtLink>
              </div>
              <div
                  @click="logout()"
                  class="text-green-500 border-green-500 border px-8 py-1 rounded-xl hover:bg-green-500 hover:text-white hover:cursor-pointer ">
                خروج
              </div>
            </div>
          </div>

          <!--          <div class="absolute inset-y-0 right-0 flex items-center pr-2 sm:static sm:inset-auto sm:ml-6 sm:pr-0">-->
          <!--            <Menu as="div" class="relative ml-3" @click="toggleSignupPopUp()" v-if="!isSignup">-->
          <!--              <div>-->


          <!--              </div>-->
          <!--            </Menu>-->
          <!--            <div v-else class="text-green-500 font-semibold"> وارد شدید ! :)</div>-->

          <!--          </div>-->
        </div>
      </div>
    </div>

<!--    <DisclosurePanel class="sm:hidden">-->
<!--      <div class="space-y-1 px-2 pt-2 pb-3">-->
<!--        <Menu as="div" class="relative ml-3" @click="toggleLoginPopUp()" v-if="!isLogin">-->
<!--          <div class="flex gap-2">-->

<!--            <div-->
<!--                class="text-green-500 border-green-500 border px-8 py-1 rounded-xl hover:cursor-pointer ">-->
<!--              ورود-->
<!--            </div>-->

<!--          </div>-->
<!--        </Menu>-->
<!--        <Menu as="div" class="relative ml-3" @click="toggleSignupPopUp()" v-if="!isSignup">-->
<!--          <div>-->
<!--            <div-->
<!--                class="text-white bg-green-500 border-green-500 border px-8 py-1 rounded-xl hover:cursor-pointer ">-->
<!--              ثبت نام-->
<!--            </div>-->
<!--          </div>-->
<!--        </Menu>-->
<!--      </div>-->
<!--    </DisclosurePanel>-->
  </Disclosure>
  <v-snackbar
      v-model="snackbar.show"
      timeout="2000"
      :color="snackbar.color"


  >
    <div class="text-center rtl">

      {{ snackbar.message }}
    </div>

    <template v-slot:actions>
      <v-btn
          variant="flat"
          @click="snackbar.show = false"
      >
        متوجه شدم
      </v-btn>
    </template>
  </v-snackbar>
  <login @successLogin="successLogin()" @togglePopUp="toggleLoginPopUp()" :is-show="isPopUpShow"/>
  <signup @successSignup="successSignup()" @signupPopup="toggleSignupPopUp()" :is-show="isSignupPopupShow"/>
</template>

<script>
import {Disclosure, DisclosureButton, DisclosurePanel, Menu, MenuButton, MenuItem, MenuItems} from '@headlessui/vue'
import {Bars3Icon, BellIcon, XMarkIcon} from '@heroicons/vue/24/outline'
import login from "./Login";

export default {
  components: {
    Bars3Icon,
    BellIcon,
    XMarkIcon,
    Disclosure,
    DisclosureButton,
    DisclosurePanel,
    Menu,
    MenuButton,
    MenuItem,
    MenuItems,
    login
  },
  created() {
    // // this.info = token;
    // // this.isLogin = this.info.isLoggedIn;
    // console.log(localStorage.hasItem('access_token'));
    // let access_token =  localStorage.hasItem('access_token') ? localStorage.getItem('access_token') : ''
    // let refresh_token =  localStorage.hasItem('refresh_token') ? localStorage.getItem('refresh_token') : ''
    // // let refresh_token = localStorage.getItem('refresh_token')
    // // let userId = localStorage.getItem('userId')
    // if (!!access_token && !!refresh_token && !!userId){
    //   this.$store.commit('setStatus', {
    //     isLoggedIn: true,
    //     access_token: access_token,
    //     refresh_token: refresh_token,
    //     Email: this.username,
    //     id: userId
    //   })
    // }
    // this.isLogin = this.$store.state.status.isLoggedIn;
  },
  mounted() {
    // this.info = token;
    // this.isLogin = this.info.isLoggedIn;
    let access_token =  localStorage.getItem('access_token');
    let refresh_token =  localStorage.getItem('refresh_token');
    let userId = localStorage.getItem('userId')

    if (!!access_token && !!refresh_token && !!userId){
      this.$store.commit('setStatus', {
        isLoggedIn: true,
        access_token: access_token,
        refresh_token: refresh_token,
        Email: this.username,
        id: userId
      })
    }
    this.isLogin = this.$store.state.status.isLoggedIn;
  },
  methods: {
    async logout() {
      let api = 'http://localhost:3000/auth/logout'

      let res = await $fetch(api, {
        method: 'POST',
      }).then(res => {
        //todo save in vueX
      }).catch(error => {
        console.log(error)
      })
      this.isLogin = false;

      this.$store.commit('setStatus', {
        isLoggedIn: false,
        access_token: '',
        refresh_token: '',
        Email: '',
      })
      this.$store.commit('setLoginData', {})
      localStorage.removeItem('access_token')
      localStorage.removeItem('refresh_token')
      localStorage.removeItem('userId')
      this.snackbar = {
        color: 'green',
        show: true,
        message: "با موفقیت خارج شدید"
      }
      this.$router.push('/')
    },
    toggleLoginPopUp() {
      this.isPopUpShow = !this.isPopUpShow;
    },
    toggleSignupPopUp() {
      this.isSignupPopupShow = !this.isSignupPopupShow;
    },
    successLogin() {
      this.isPopUpShow = !this.isPopUpShow;
      this.snackbar = {
        color: 'green',
        show: true,
        message: "با موفقیت وارد شدید"
      }
      localStorage.setItem('access_token',this.$store.state.status.access_token)
      localStorage.setItem('refresh_token',this.$store.state.status.refresh_token)
      localStorage.setItem('userId',this.$store.state.status.id)
      this.isLogin = true;
    },
    successSignup() {
      this.isSignupPopupShow = !this.isSignupPopupShow;
      this.isSignup = true;
      this.snackbar = {
        color: 'green',
        show: true,
        message: "با موفقیت ثبت نام شدید"
      }
    },
    toggleNavigationActivation(item) {
      this.navigation.map(nav => {
        nav.current = nav.name === item.name;
      })
    }
  },
  data() {
    return {
      isPopUpShow: false,
      isSignupPopupShow: false,
      isLogin: false,

      snackbar: {
        color: null,
        show: false,
        message: "با موفقیت پاک شد"
      },
      info: {
        isLoggedIn: false,
        access_token: '',
        refresh_token: '',
        Email: '',
        name: ''
      },

      isSignup: false,

      navigation: [
        // {name: 'خانه', href: '/', current: false},
        // {name: 'محصولات', href: '#', current: false},
        // {name: 'خدمات', href: '#', current: false},
        // {name: 'وبلاگ', href: '/blogs', current: false},
      ]
    }
  },
}

</script>
