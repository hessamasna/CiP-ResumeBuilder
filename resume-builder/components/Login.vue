<template>
  <div v-if="isShow"
       class="min-w-screen h-screen animated fadeIn faster  fixed  left-0 top-0 flex justify-center items-center inset-0 z-50 outline-none focus:outline-none bg-no-repeat bg-center bg-cover"
       id="modal-id">
    <div @click="togglePopUp()" class="absolute bg-black opacity-80 inset-0 z-0"></div>
    <div class="w-full  max-w-5xl  p-5 relative mx-auto my-auto rounded-xl shadow-lg  bg-white ">
      <!--content-->
      <div class="">
        <!--body-->
        <div class="p-5 grid w-full max-w-sm mx-auto">
          <h2 class="text-4xl font-bold my-4 mb-10 text-center">ورود به حساب کاربری</h2>
          <div class="mt-2">
            <label for="email-address"
                   :class=" ['text-sm font-medium text-gray-700',this.showUsernameError?'text-red-600':'']">نام
              کاربری</label>
            <div class="relative mt-2 max-w-96">
              <div class="absolute inset-y-0 right-0 flex items-center pr-3 pointer-events-none">
                <svg width="21" height="21" viewBox="0 0 21 21" fill="none" xmlns="http://www.w3.org/2000/svg">
                  <path fill-rule="evenodd" clip-rule="evenodd"
                        d="M10.0001 1.75C8.04735 1.75 6.46436 3.33299 6.46436 5.28571C6.46436 7.23844 8.04735 8.82143 10.0001 8.82143C11.9528 8.82143 13.5358 7.23844 13.5358 5.28571C13.5358 3.33299 11.9528 1.75 10.0001 1.75ZM4.96436 5.28571C4.96436 2.50457 7.21892 0.25 10.0001 0.25C12.7812 0.25 15.0358 2.50457 15.0358 5.28571C15.0358 8.06686 12.7812 10.3214 10.0001 10.3214C7.21892 10.3214 4.96436 8.06686 4.96436 5.28571Z"
                        fill="black" fill-opacity="0.2"/>
                  <path fill-rule="evenodd" clip-rule="evenodd"
                        d="M2.31007 15.1661C3.86731 13.5023 6.39993 12.25 10.2862 12.25C14.1725 12.25 16.7051 13.5023 18.2623 15.1661C19.8 16.8088 20.3219 18.7847 20.3219 20.1429C20.3219 20.5571 19.9861 20.8929 19.5719 20.8929H1.00049C0.586275 20.8929 0.250488 20.5571 0.250488 20.1429C0.250488 18.7847 0.772444 16.8088 2.31007 15.1661ZM1.81427 19.3929H18.7581C18.6021 18.4301 18.1405 17.2309 17.1672 16.1911C15.9387 14.8786 13.8285 13.75 10.2862 13.75C6.74391 13.75 4.63367 14.8786 3.40519 16.1911C2.43191 17.2309 1.97034 18.4301 1.81427 19.3929Z"
                        fill="black" fill-opacity="0.2"/>
                </svg>
              </div>
              <input type="text" id="input-group-1" v-model="username"
                     :class="['border border-solid border-gray-300 rounded-xl w-full ps-10 text-sm p-3 h-12', this.showUsernameError?'border-red-600 text-red-600':'']"
                     placeholder="نام کاربری خود را وارد کنید">
            </div>
          </div>
          <div class="mt-3">
            <label for="email-address"
                   :class=" ['text-sm font-medium text-gray-700',this.showPasswordError?'text-red-600':'']">رمز
              عبور</label>
            <div class="relative mt-2 max-w-96">
              <div class="absolute inset-y-0 right-0 flex items-center pr-3 pointer-events-none">
                <svg width="18" height="22" viewBox="0 0 18 22" fill="none" xmlns="http://www.w3.org/2000/svg">
                  <path fill-rule="evenodd" clip-rule="evenodd"
                        d="M0.650391 13.3334C0.650391 11.2623 2.32932 9.58337 4.40039 9.58337H13.6004C15.6715 9.58337 17.3504 11.2623 17.3504 13.3334V15.6667C17.3504 19.0264 14.6268 21.75 11.2671 21.75H6.73372C3.37399 21.75 0.650391 19.0264 0.650391 15.6667V13.3334ZM4.40039 11.0834C3.15775 11.0834 2.15039 12.0907 2.15039 13.3334V15.6667C2.15039 18.198 4.20242 20.25 6.73372 20.25H11.2671C13.7984 20.25 15.8504 18.198 15.8504 15.6667V13.3334C15.8504 12.0907 14.843 11.0834 13.6004 11.0834H4.40039Z"
                        fill="black" fill-opacity="0.2"/>
                  <path fill-rule="evenodd" clip-rule="evenodd"
                        d="M8.99998 17.9001C8.50292 17.9001 8.09998 17.4972 8.09998 17.0001L8.09998 14.3335C8.09998 13.8364 8.50292 13.4335 8.99998 13.4335C9.49703 13.4335 9.89998 13.8364 9.89998 14.3335L9.89998 17.0001C9.89998 17.4972 9.49703 17.9001 8.99998 17.9001Z"
                        fill="black" fill-opacity="0.2"/>
                  <path fill-rule="evenodd" clip-rule="evenodd"
                        d="M6.9715 2.59587C6.46405 3.02841 6.10017 3.68694 6.10017 4.64389V9.66664C6.10017 10.1637 5.69723 10.5666 5.20017 10.5666C4.70311 10.5666 4.30017 10.1637 4.30017 9.66664V4.64389C4.30017 3.17152 4.88629 2.00808 5.80384 1.22599C6.70211 0.460328 7.86835 0.099962 9.00018 0.0999756C10.132 0.0999892 11.2982 0.460383 12.1965 1.22605C13.114 2.00815 13.7002 3.17158 13.7002 4.64392V9.66664C13.7002 10.1637 13.2972 10.5666 12.8002 10.5666C12.3031 10.5666 11.9002 10.1637 11.9002 9.66664V4.64392C11.9002 3.68701 11.5363 3.02848 11.0288 2.59592C10.5021 2.14694 9.76833 1.89999 9.00016 1.89998C8.23199 1.89997 7.49823 2.1469 6.9715 2.59587Z"
                        fill="black" fill-opacity="0.2"/>
                </svg>

              </div>
              <input :type="showPassword? 'text': 'password'" id="input-group-1" v-model="password"
                     :class="['border border-solid border-gray-300 rounded-xl w-full ps-10 text-sm p-3 h-12', this.showPasswordError?'border-red-600 text-red-600':'']"
                     placeholder="رمز عبور خود را وارد کنید">

              <div class="absolute inset-y-0 left-0 flex items-center pl-3 hover:cursor-pointer "
                   @click="showPassword = !showPassword">

                <svg width="22" height="14" viewBox="0 0 22 14" fill="none" xmlns="http://www.w3.org/2000/svg" class=""
                >
                  <path fill-rule="evenodd" clip-rule="evenodd"
                        d="M2.03439 7.11743C2.54874 7.70671 3.54497 8.77022 4.84692 9.78946C6.5985 11.1607 8.79663 12.3674 11.0566 12.3674C13.3165 12.3674 15.5146 11.1607 17.2662 9.78946C18.5682 8.77022 19.5644 7.70671 20.0788 7.11743C19.5644 6.52815 18.5682 5.46465 17.2662 4.4454C15.5146 3.07417 13.3165 1.86743 11.0566 1.86743C8.79663 1.86743 6.59849 3.07417 4.84692 4.4454C3.54497 5.46465 2.54874 6.52815 2.03439 7.11743ZM3.92228 3.26429C5.76664 1.82042 8.29474 0.367432 11.0566 0.367432C13.8184 0.367432 16.3465 1.82042 18.1909 3.26429C20.0521 4.72136 21.3258 6.25525 21.5937 6.58797C21.8436 6.89822 21.8436 7.33663 21.5937 7.64689C21.3258 7.97961 20.0521 9.5135 18.1909 10.9706C16.3465 12.4144 13.8184 13.8674 11.0566 13.8674C8.29474 13.8674 5.76664 12.4144 3.92228 10.9706C2.06105 9.5135 0.787386 7.97961 0.519429 7.64689C0.269571 7.33665 0.26956 6.89823 0.519429 6.58797C0.787384 6.25525 2.06105 4.72137 3.92228 3.26429Z"
                        :fill="showPassword? '#00C853': 'rgba(0, 0, 0, 0.2)'"/>
                  <path fill-rule="evenodd" clip-rule="evenodd"
                        d="M11.0566 4.86743C9.814 4.86743 8.80664 5.87479 8.80664 7.11743C8.80664 8.36007 9.814 9.36743 11.0566 9.36743C12.2993 9.36743 13.3066 8.36007 13.3066 7.11743C13.3066 5.87479 12.2993 4.86743 11.0566 4.86743ZM7.30664 7.11743C7.30664 5.04636 8.98557 3.36743 11.0566 3.36743C13.1277 3.36743 14.8066 5.04636 14.8066 7.11743C14.8066 9.1885 13.1277 10.8674 11.0566 10.8674C8.98557 10.8674 7.30664 9.1885 7.30664 7.11743Z"
                        :fill="showPassword? '#00C853': 'rgba(0, 0, 0, 0.2)'"/>
                </svg>
              </div>

            </div>
          </div>
          <button
              class=" mt-10 rounded-xl text-white align-middle bg-green-500 py-2" @click="login()">
            ورود
          </button>
        </div>
        <!--footer-->
        <div class="p-3  mt-2 text-center space-x-4 md:block">

        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: "Login",
  props: ['isShow'],
  methods: {
    togglePopUp() {
      this.$emit('togglePopUp')
    },
    async login() {
      if (this.username.length === 0) {
        this.showUsernameError = true
        return;
      }

      if (this.password.length === 0) {
        this.showPasswordError = true
        return;
      }

      let api = 'https://challenge.webjar.ir/auth/login'
      let body = {
        userName: this.username,
        password: this.password
      }

      let res = await $fetch(api, {
        method: 'POST',
        body: JSON.stringify(body)
      }).then(res => {
        //todo save in vueX
        this.showPasswordError = false;
        this.showUsernameError = false;
        this.$emit('successLogin')
      }).catch(error => {
        this.showPasswordError = true;
        this.showUsernameError = true;
        console.log(error)
      })

    }
  },
  data() {
    return {
      showUsernameError: false,
      showPasswordError: false,
      showPassword: false,
      username: '',
      password: ''
    }
  },
}
</script>

<style scoped>

</style>