<template>
  <loading class="h-screen" v-if="loading"/>
  <div class="pa-8" v-else>
<!--    <div class="text-center text-2xl flex flex-column gap-2" >-->
<!--      <span> رزومه فعالی ندارید :)</span>-->
<!--      <v-btn color="green" to="new-resume">-->
<!--        ساخت رزومه جدید-->
<!--      </v-btn>-->
<!--    </div>-->

    <div class="grid grid-flow-col auto-cols-max gap-8 	">

      <v-card
          v-for="(resume,index) in resumes"
          :key="resumes.id? resumes.id: index"
          class="mx-auto"
          max-width="344"
      >
        <div class="" style="height: 150px">
          <v-img
              :src="`/img/Cv_${resume.template_number}_preview.png`"
              fit
          ></v-img>
          <!--<span class="absolute inset-16">  تملیت یک</span>-->
        </div>

        <v-card-title>
          {{ resume.title }}
        </v-card-title>

        <v-card-subtitle>
          {{ resume.personal_info.first_name }} {{ resume.personal_info.last_name }} - {{ resume.job_title }}
        </v-card-subtitle>

        <v-card-actions>
          <NuxtLink :to="`Cv-${resume.template_number}-${resume.id}`" class="">
            <v-btn
                color="green"
                variant="text"
            >
              نمایش
            </v-btn>
          </NuxtLink>
          <v-spacer></v-spacer>
          <NuxtLink :to="`new-resume-${resume.id}`" class="">
            <v-btn
                class="px-0 mx-0"
                color="orange-lighten-2"
            >
              ویرایش
            </v-btn>
          </NuxtLink>

          <v-btn
              class="px-0 mx-0"
              color="red"
              @click="deleteCv(resume.id)"
          >
            حذف
          </v-btn>


        </v-card-actions>


      </v-card>

      <NuxtLink  to="new-resume-0" class="bg-green border-dashed border-4 border-black w-52 flex justify-center align-center h-[270px]">
          <v-icon icon="mdi-plus" size="large"></v-icon>
          ساخت رزومه جدید
      </NuxtLink>

    </div>
    <v-snackbar
        v-model="snackbar.show"
        :timeout="5000"
        :color="snackbar.color"
    >
      {{ snackbar.message }}

      <template v-slot:actions>
        <v-btn
            variant="flat"
            @click="snackbar.show = false"
        >
          متوجه شدم
        </v-btn>
      </template>
    </v-snackbar>
  </div>
</template>

<script>
export default {
  name: "dashboard",
  async created() {
    if (!this.$store.state.status.isLoggedIn) {
      this.$router.push('/');
    } else {
      this.loading = true;
      let api = 'http://localhost:3000/cv/getAll/' + this.$store.state.status.id;
      // console.log(this.$store.state.cookie)
      let res = await $fetch(api, {
        method: 'GET',
        headers: {
          'access_token': this.$store.state.status.access_token,
          'refresh_token': this.$store.state.status.refresh_token
        },
      }).then(res => {
        console.log(res)
        this.resumes = res.result;
      }).catch(error => {
        console.log(error)
      })
    }

    this.loading = false;

  },
  methods: {
    async deleteCv(id) {
      this.loading = true;

      let api = 'http://localhost:3000/cv/delete/' + id;

      let res = await $fetch(api, {
        method: 'DELETE',
        headers: {
          'access_token': this.$store.state.status.access_token,
          'refresh_token': this.$store.state.status.refresh_token
        },
      }).then(res => {
        this.snackbar.message = "با موفقیت پاک شد"
        this.snackbar.color = "green"
        this.snackbar.show = true;
      }).catch(error => {
        this.snackbar.message = "مشکل در پاک کردن"
        this.snackbar.color = 'red'
        this.snackbar.show = true;
        console.log(error)
      })
      this.resumes = this.resumes.filter(object => {
        return object.id !== id;
      });

      this.loading = false;

    },

    previewCv(id) {

    }
  },
  data() {
    return {
      loading: false,
      snackbar: {
        color: null,
        show: false,
        message: "با موفقیت پاک شد"
      },
      resumes: []
    }
  },
}
</script>

<style scoped>

</style>
