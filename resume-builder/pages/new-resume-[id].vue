<template>
  <div class="p-12">
    <div class="font-bold mb-10 text-2xl">ساخت رزومه جدید</div>
    <v-expansion-panels>
      <v-expansion-panel class="pa-5">
        <v-expansion-panel-title>اطلاعات کلی</v-expansion-panel-title>
        <v-expansion-panel-text class="mt-5">
          <v-row>
            <v-col cols="12" sm="4">
              <v-text-field label="عنوان رزومه" v-model="resume.title"/>
            </v-col>
          </v-row>
          <v-row>
            <v-col cols="12" sm="4">
              <v-text-field label="عنوان شغلی" v-model="resume.job_title"/>
            </v-col>
          </v-row>
          <v-row>
            <v-col cols='12' sm="4">
              <v-select :items="resumeTemplates" item-title="description" item-value="id" label="قالب مدنظر شما"
                        v-model="resume.template_number"/>
            </v-col>
          </v-row>

        </v-expansion-panel-text>
      </v-expansion-panel>
      <v-expansion-panel class="pa-5" :title="FORM_NAMES.PERSONAL_INFO">
        <v-expansion-panel-text class="mt-4">
          <v-row>
            <v-col cols="12" sm='4'>
              <v-text-field :label="FORM_NAMES.FIRST_NAME" v-model="resume.personal_info.first_name"/>
            </v-col>
            <v-col cols="12" sm='4'>
              <v-text-field :label="FORM_NAMES.LAST_NAME" v-model="resume.personal_info.last_name"/>
            </v-col>
            <v-col cols="12" sm='4'>
              <v-text-field :label="FORM_NAMES.AGE" v-model="resume.personal_info.age"/>
            </v-col>
          </v-row>
          <v-row>
            <v-col cols="12" sm="4">
              <v-text-field :label="FORM_NAMES.PHONE" v-model="resume.personal_info.phone_number" counter
                            maxlength="11"/>
            </v-col>
            <v-col cols="12" sm="4">
              <v-text-field :label="FORM_NAMES.EMAIL" v-model="resume.personal_info.email"/>
            </v-col>
            <v-col cols="12" sm="4">
              <v-text-field :label="FORM_NAMES.LOCATION" v-model="resume.personal_info.address"/>
            </v-col>
            <!--            <v-col cols="12" sm="4">-->
            <!--              <v-file-input label="عکس پروفایل" v-model="resume.image"/>-->
            <!--            </v-col>-->
          </v-row>
          <v-row>
            <v-col cols="12">
              <v-textarea :label="FORM_NAMES.ABOUT_ME" v-model="resume.about_me" clearable
                          :hint="FORM_NAMES.ABOUT_ME_PLACEHOLDER" counter/>
            </v-col>
          </v-row>
        </v-expansion-panel-text>
      </v-expansion-panel>
      <v-expansion-panel class="pa-5" :title="FORM_NAMES.EDUCATIONAL_INFO">
        <v-expansion-panel-text class="mt-4">
          <v-row v-for="(education,index) in resume.education" :key="index" class="border-b">
            <v-col cols=12 sm="3">
              <v-select :items="educationGrades" label="مقطع تحصیلی" v-model="education.degree"/>
            </v-col>
            <v-col cols=12 sm="3">
              <v-text-field label="نام دانشگاه" v-model="education.school"/>
            </v-col>
            <v-col cols=12 sm="3">
              <v-text-field label="سال آغار" v-model="education.start"/>
            </v-col>
            <v-col cols=12 sm="3">
              <v-text-field label="سال پایان" v-model="education.end"/>
            </v-col>
          </v-row>
          <v-row class="mt-5">
            <v-col cols=12 sm=3>
              <v-btn variant="outlined" @click="addNewEducationalHistory">
                اضافه سابقه جدید
              </v-btn>
            </v-col>
          </v-row>
        </v-expansion-panel-text>
      </v-expansion-panel>
      <v-expansion-panel title="سوابق کاری" class="pa-5">
        <v-expansion-panel-text>
          <div v-for="(workExperience,index) in resume.experience" :key="index" class="border-b">
            <v-row class="mt-5">
              <v-col cols=12 sm="3">
                <v-text-field label="عنوان شغلی" v-model="workExperience.title"/>
              </v-col>
            </v-row>
            <v-row>
              <v-col cols="12" sm="4">
                <v-text-field label="نام شرکت" v-model="workExperience.company"/>
              </v-col>
              <v-col cols="6" sm="2">
                <v-text-field label="سال شروع" v-model="workExperience.start"/>
              </v-col>
              <v-col cols="6" sm="2">
                <v-text-field label="سال پایان" v-model="workExperience.end"/>
              </v-col>
            </v-row>
            <v-row>
              <v-col cols="12">
                <v-textarea label="درباره سابقه" v-model="workExperience.description" clearable counter/>
              </v-col>
            </v-row>
          </div>
          <v-row class="mt-5">
            <v-col cols=12 sm=3>
              <v-btn variant="outlined" @click="addNewWorkExperience">
                اضافه سابقه جدید
              </v-btn>
            </v-col>
          </v-row>
        </v-expansion-panel-text>
      </v-expansion-panel>
      <v-expansion-panel class="pa-5">
        <v-expansion-panel-title>تکنولوژی ها</v-expansion-panel-title>
        <v-expansion-panel-text class="mt-5">
          <v-row v-for="(skill,index) in resume.skills" :key="index" class="border-b mt-5">
            <v-col cols=12 sm="3">
              <v-text-field label="توانایی" v-model="skill.name"/>
            </v-col>
            <v-col cols=12 sm="3">
              <v-text-field label="میزان تسلط" v-model.number="skill.percent" hint="در بازه ۱ تا ۱۰۰"/>
            </v-col>
          </v-row>
          <v-row>
            <v-col cols=12 sm=3>
              <v-btn variant="outlined" @click="addNewSkill">
                اضافه توانایی جدید
              </v-btn>
            </v-col>
          </v-row>
        </v-expansion-panel-text>
      </v-expansion-panel>
      <v-expansion-panel class="pa-5">
        <v-expansion-panel-title>راه های ارتباطی</v-expansion-panel-title>
        <v-expansion-panel-text class="mt-5">
          <v-row v-for="(media,index) in resume.social_medias" :key="index" class="border-b">
            <v-col cols="6" sm="3">
              <v-select :items="socialMedias" item-value="value" item-text="title" label="راه ارتباطی"
                        v-model="media.plat_form"/>
            </v-col>
            <v-col cols="6" sm="3">
              <v-text-field label="آدرس" v-model="media.link"/>
            </v-col>
          </v-row>
          <v-row>
            <v-col cols=12 sm=3>
              <v-btn variant="outlined" @click="addNewSocialMedia">
                اضافه راه ارتباطی جدید
              </v-btn>
            </v-col>
          </v-row>
        </v-expansion-panel-text>
      </v-expansion-panel>
    </v-expansion-panels>
    <v-btn @click="submit" color="green" variant="flat" block class="text-center mt-8 px-8">ذخیره</v-btn>
    <v-snackbar
        v-model="snackbar.show"
        :timeout="5000"
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
  </div>

</template>

<script>
export default {
  async created() {
    let id = this.$route.params.id
    if (id != 0) {
      let api = 'http://localhost:3000/cv/get/' + id;

      let res = await $fetch(api, {
        method: 'GET',
        headers: {
          'access_token': this.$store.state.status.access_token,
          'refresh_token': this.$store.state.status.refresh_token
        },
      }).then(res => {
        // //todo
        // this.resume.font_size = res.font_size
        // this.resume.font_family = res.font_family
        // this.resume.color = res.color
        // this.resume.personal_info = res.personal_info
        // this.resume.template_number = res.template_number
        // this.resume.job_title = res.job_title
        // this.resume.location = res.location
        // this.resume.image = res.image
        // this.resume.education = res.about_me
        // this.resume = {...res.result}
        this.resume = res.result
        this.resume.personal_info.age = 22
        // this.resume.image = res.result.image

        this.loading = false;

      }).catch(error => {
        this.loading = true;
        this.loading = false;
        console.log(error)
      })
    }
  },
  data() {
    return {
      snackbar: {
        color: null,
        show: false,
        message: "با موفقیت پاک شد"
      },
      FORM_NAMES: {
        PERSONAL_INFO: 'اطلاعات شخصی',
        EDUCATIONAL_INFO: 'اطلاعات تحصیلی',
        EMAIL: 'ایمیل',
        PHONE: 'شماره همراه',
        AGE: 'سن',
        LOCATION: 'محل سکونت',
        FIRST_NAME: 'نام',
        LAST_NAME: 'نام خانوادگی',
        ABOUT_ME: 'درباره من',
        ABOUT_ME_PLACEHOLDER: 'چند خطی درباره من',
      },
      resume: {
        font_size: 15,
        font_family: 'IRANSans',
        color: '#FFA500',
        personal_info: {
          first_name: '',
          last_name: '',
          email: '',
          phone_number: '',
          address: '',
          age: 22
        },
        title: '',
        template_number: 1,
        job_title: '',
        location: '',
        is_public: true,
        image: '',
        about_me: '',
        education: [{
          degree: '',
          major: '', //TODO ADD FORM FIELD FOR THIS
          school: '',
          start: '',
          end: ''
        }],
        experience: [{
          title: '',
          company: '',
          start: '',
          end: '',
          description: ''
        }],
        skills: [{
          name: '',
          percent: 0
        }],
        social_medias: [{
          plat_form: '',
          link: ''
        }]
      },
      validators: {
        required: value => !!value || 'Required.'
      },
      educationGrades: ['کارشناسی', 'کارشناسی ارشد', 'دکتری'],
      socialMedias: [
        {
          title: 'لینکدین',
          value: 'linkedin'
        }, {
          title: 'اینستاگرام',
          value: 'instagram'
        }, {
          title: 'گیت هاب',
          value: 'github'
        },
      ],
      resumeTemplates: [
        {
          id: 1,
          description: 'قالب اول'
        },
        {
          id: 2,
          description: 'قالب دوم'
        },
        {
          id: 4,
          description: 'قالب سوم'
        }
      ]
    }
  },
  methods: {
    addNewEducationalHistory() {
      this.resume.education.push(
          {
            degree: '',
            school: '',
            start: '',
            end: ''
          }
      )
    },
    addNewWorkExperience() {
      this.resume.experience.push({
        title: '',
        company: '',
        start: '',
        end: '',
        description: ''
      })
    },
    addNewSkill() {
      this.resume.skills.push({
        name: '',
        percent: ''
      })
    },
    addNewSocialMedia() {
      this.resume.social_medias.push({
        plat_form: '',
        link: ''
      })
    },
    async submit() {
      let id = this.$route.params.id
      if (id != 0) {
        console.log("y")

        let api = 'http://localhost:3000/cv/update';
        let res = await $fetch(api, {
          method: 'PUT',
          body: JSON.stringify(this.resume),
          headers: {
            'access_token': this.$store.state.status.access_token,
            'refresh_token': this.$store.state.status.refresh_token
          },
        }).then(res => {
          console.log(res)
          this.snackbar = {
            color: 'green',
            show: true,
            message: 'با موفقیت ذخیره شد'
          }
          this.$router.push('/dashboard');
        }).catch(error => {
          this.snackbar = {
            color: 'red',
            show: true,
            message: 'خطایی رخ داده است'
          }
          console.log(error)
        })
      } else {
        const y = JSON.stringify(this.resume)
        console.log('2')
        // let fd = new FormData()
        // fd.append(y)
        console.log(y)
        let api = 'http://localhost:3000/cv/create';
        let res = await $fetch(api, {
          method: 'POST',
          body: JSON.stringify(this.resume),
          headers: {
            'access_token': this.$store.state.status.access_token,
            'refresh_token': this.$store.state.status.refresh_token
          },
        }).then(res => {
          this.snackbar = {
            color: 'green',
            show: true,
            message: 'با موفقیت ذخیره شد'
          }
          this.$router.push('/dashboard');
          console.log(res)
        }).catch(error => {
          this.snackbar = {
            color: 'red',
            show: true,
            message: 'خطایی رخ داده است'
          }
          console.log(error)
        })
      }
    }
  }
}
</script>

<style>

</style>
