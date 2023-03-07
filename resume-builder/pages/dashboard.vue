<template>
  <div class="pa-8">
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
          {{ resume.name }}
        </v-card-title>

        <v-card-subtitle>
          {{ resume.personal_info.first_name }} {{ resume.personal_info.last_name }} -
          <!--          {{ resume.personal.cvTitle }}-->
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
          <NuxtLink :to="`edit/Cv-${resume.template_number}-${resume.id}`" class="">
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
    console.log('isLoggedIn: ' + this.$store.state.status.isLoggedIn);
    if (!this.$store.state.status.isLoggedIn) {
      this.$router.push('/');
    } else {
      let api = 'http://localhost:3000/cv/getAll/' + this.$store.state.status.id;
      console.log(this.$store.state.cookie)
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
  },
  methods: {
    async deleteCv(id) {
      let api = 'https://localhost:3000/cvDelete/' + id;

      let res = await $fetch(api, {
        method: 'DELETE',
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
    },
    previewCv(id) {

    }
  },
  data() {
    return {
      snackbar: {
        color: null,
        show: false,
        message: "با موفقیت پاک شد"
      },
      resumes:
          [
            {
              "id": 1,
              "user_id": 2,
              "personal_info": {
                "id": 1,
                "first_name": "John",
                "last_name": "Doe",
                "email": "john.doe@example.com",
                "phone_number": "+1-555-555-5555",
                "address": "123 Main St, Anytown USA"
              },
              "about_me": "",
              "education": [],
              "experience": [],
              "skills": [
                {
                  "id": 1,
                  "cv_id": 1,
                  "name": "Java",
                  "percent": 90
                },
                {
                  "id": 2,
                  "cv_id": 1,
                  "name": "Python",
                  "percent": 85
                },
                {
                  "id": 3,
                  "cv_id": 1,
                  "name": "JavaScript",
                  "percent": 80
                },
                {
                  "id": 4,
                  "cv_id": 1,
                  "name": "SQL",
                  "percent": 75
                }
              ],
              "social_medias": [
                {
                  "id": 1,
                  "cv_id": 1,
                  "plat_form": "",
                  "link": "https://www.linkedin.com/in/johndoe/"
                },
                {
                  "id": 2,
                  "cv_id": 1,
                  "plat_form": "",
                  "link": "https://github.com/johndoe"
                }
              ],
              "name": "John Doe",
              "is_public": false,
              "font_size": 12,
              "font_family": "Arial",
              "color": "#333",
              "template_number": 1
            },
            {
              "id": 2,
              "user_id": 2,
              "personal_info": {
                "id": 1,
                "first_name": "John1",
                "last_name": "Doe2",
                "email": "john.doe@example.com",
                "phone_number": "+1-555-555-5555",
                "address": "123 Main St, Anytown USA"
              },
              "about_me": "",
              "education": [],
              "experience": [],
              "skills": [
                {
                  "id": 5,
                  "cv_id": 2,
                  "name": "Java",
                  "percent": 90
                },
                {
                  "id": 6,
                  "cv_id": 2,
                  "name": "Python",
                  "percent": 85
                },
                {
                  "id": 7,
                  "cv_id": 2,
                  "name": "JavaScript",
                  "percent": 80
                },
                {
                  "id": 8,
                  "cv_id": 2,
                  "name": "SQL",
                  "percent": 75
                }
              ],
              "social_medias": [
                {
                  "id": 3,
                  "cv_id": 2,
                  "plat_form": "",
                  "link": "https://www.linkedin.com/in/johndoe/"
                },
                {
                  "id": 4,
                  "cv_id": 2,
                  "plat_form": "",
                  "link": "https://github.com/johndoe"
                }
              ],
              "name": "John Doe",
              "is_public": false,
              "font_size": 12,
              "font_family": "Arial",
              "color": "#333",
              "template_number": 1
            }

          ]
    }
  },
}
</script>

<style scoped>

</style>
