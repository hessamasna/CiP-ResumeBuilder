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
              :src="`/img/Cv_${resume.cv_temp}_preview.png`"
              fit
          ></v-img>
<!--<span class="absolute inset-16">  تملیت یک</span>-->
        </div>

        <v-card-title>
          {{ resume.cv_name }}
        </v-card-title>

        <v-card-subtitle>
          {{ resume.personal.name }} - {{ resume.personal.cvTitle }}
        </v-card-subtitle>

        <v-card-actions>
<!--          <v-btn-->
<!--              color="green"-->
<!--              variant="text"-->
<!--              :to="`Cv-${resumes.cv_temp}-${resumes.id}`"-->
<!--          >-->
<!--            نمایش-->
<!--          </v-btn>-->
          <NuxtLink :to="`Cv-${resume.cv_temp}-${resume.id}`" class="">
            <v-btn
                color="green"
                variant="text"
            >
              نمایش
            </v-btn>
          </NuxtLink>
          <v-spacer></v-spacer>
          <NuxtLink :to="`edit/Cv-${resume.cv_temp}-${resume.id}`" class="">
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
    previewCv(id){

    }
  },
  data() {
    return {
      snackbar: {
        color: null,
        show: false,
        message: "با موفقیت پاک شد"
      },
      resumes: [
        {
          id:2,
          image: "https://tailone.tailwindtemplate.net/src/img/dummy/avatar1.png",
          font: 'vazir-FD',
          color: "#ffa500",
          fontSize: 15,
          cv_temp: 1,
          cv_name: "رزومه اول",
          personal: {
            name: "سیپ سی ویq",
            cvTitle: "توسعه دهنده",
            phone: "09121112233",
            age:22,
            email: "Cv@cipCv.com",
            website: "www.cipCv.com",
            address: "Tehran",
            social: [
              {
                title: "linkedin",
                path: "CipCV",
                url: "https://linkedin.com/cipC"
              },
              {
                title: "instagram",
                path: "CipCV",
                url: "https://instagram.com/CipCv"
              },
            ]
          },
          aboutMe: "لورم ایپسوم متن ساختگی با تولید سادگی نامفهوم از صنعت چاپ، و با استفاده از طراحان گرافیک است، چاپگرها و متون بلکه روزنامه و مجله در ستون و سطرآنچنان که لازم است، و برای شرایط فعلی تکنولوژی مورد نیاز، و کاربردهای متنوع با هدف بهبود ابزارهای کاربردی می باشد، کتابهای زیادی در شصت و سه درصد گذشته حال و آینده، شناخت فراوان جامعه و متخصصان را می طلبد، تا با نرم افزارها شناخت بیشتری را برای طراحان رایانه ای علی الخصوص طراحان خلاقی، و فرهنگ پیشرو در زبان فارسی ایجاد کرد، در این صورت می توان امید داشت که تمام و دشواری موجود در ارائه راهکارها، و شرایط سخت تایپ به پایان رسد و زمان مورد نیاز شامل حروفچینی دستاوردهای اصلی، و جوابگوی سوالات پیوسته اهل دنیای موجود طراحی اساسا مورد استفاده قرار گیرد.",
          workHistory: [
            {
              title: 'FrontEnd Developer',
              startDate: '2021',
              endDate: '2022',
              place: "Sharif",
              description: "Lorem ipsum dolor sit amet, consectetur adipisicing elit. Dolorem explicabo impedit magni reiciendis suscipit! Adipisci deleniti earum eveniet incidunt ipsa libero modi, molestiae nobis pariatur quod repellat, repudiandae rerum sit!"
            },
            {
              title: 'توسعه دهنده فرانت',
              startDate: '1399',
              endDate: '1400',
              place: "دانشگاه صنعتی شریف",
              description: " زمان مورد نیاز شامل حروفچینی دستاوردهای اصلی، و جوابگوی سوالات پیوسته اهل دنیای موجود طراحی اساسا مورد استفاده قرار گیرد."
            }
          ],
          skills: [
            {
              title: "Java",
              amount: 80,
            },
            {
              title: 'JS',
              amount: 85,
            },
            {
              title: 'Ruby',
              amount: 20,
            }
          ],
          educations: [
            {
              title: "مهندسی کامپیوتر",
              place: "صنعتی شریف",
              degree: "کارشناسی",
              startDate: "1398",
              endDate: "1402",
            },
            {
              title: "مهندسی کامپیوتر",
              place: "صنعتی شریف",
              degree: "کارشناسی ارشد",
              startDate: "1402",
              endDate: "1406",
            },
          ]
        },
        {
          id:1,
          image: "https://tailone.tailwindtemplate.net/src/img/dummy/avatar1.png",
          font: 'vazir-FD',
          color: "#ffa500",
          fontSize: 15,
          cv_temp: 2,
          cv_name: "رزومه دوم",
          personal: {
            name: "سیپ سی ویq",
            cvTitle: "توسعه دهنده",
            phone: "09121112233",
            age:22,
            email: "Cv@cipCv.com",
            website: "www.cipCv.com",
            address: "Tehran",
            social: [
              {
                title: "linkedin",
                path: "CipCV",
                url: "https://linkedin.com/cipC"
              },
              {
                title: "instagram",
                path: "CipCV",
                url: "https://instagram.com/CipCv"
              },
            ]
          },
          aboutMe: "لورم ایپسوم متن ساختگی با تولید سادگی نامفهوم از صنعت چاپ، و با استفاده از طراحان گرافیک است، چاپگرها و متون بلکه روزنامه و مجله در ستون و سطرآنچنان که لازم است، و برای شرایط فعلی تکنولوژی مورد نیاز، و کاربردهای متنوع با هدف بهبود ابزارهای کاربردی می باشد، کتابهای زیادی در شصت و سه درصد گذشته حال و آینده، شناخت فراوان جامعه و متخصصان را می طلبد، تا با نرم افزارها شناخت بیشتری را برای طراحان رایانه ای علی الخصوص طراحان خلاقی، و فرهنگ پیشرو در زبان فارسی ایجاد کرد، در این صورت می توان امید داشت که تمام و دشواری موجود در ارائه راهکارها، و شرایط سخت تایپ به پایان رسد و زمان مورد نیاز شامل حروفچینی دستاوردهای اصلی، و جوابگوی سوالات پیوسته اهل دنیای موجود طراحی اساسا مورد استفاده قرار گیرد.",
          workHistory: [
            {
              title: 'FrontEnd Developer',
              startDate: '2021',
              endDate: '2022',
              place: "Sharif",
              description: "Lorem ipsum dolor sit amet, consectetur adipisicing elit. Dolorem explicabo impedit magni reiciendis suscipit! Adipisci deleniti earum eveniet incidunt ipsa libero modi, molestiae nobis pariatur quod repellat, repudiandae rerum sit!"
            },
            {
              title: 'توسعه دهنده فرانت',
              startDate: '1399',
              endDate: '1400',
              place: "دانشگاه صنعتی شریف",
              description: " زمان مورد نیاز شامل حروفچینی دستاوردهای اصلی، و جوابگوی سوالات پیوسته اهل دنیای موجود طراحی اساسا مورد استفاده قرار گیرد."
            }
          ],
          skills: [
            {
              title: "Java",
              amount: 80,
            },
            {
              title: 'JS',
              amount: 85,
            },
            {
              title: 'Ruby',
              amount: 20,
            }
          ],
          educations: [
            {
              title: "مهندسی کامپیوتر",
              place: "صنعتی شریف",
              degree: "کارشناسی",
              startDate: "1398",
              endDate: "1402",
            },
            {
              title: "مهندسی کامپیوتر",
              place: "صنعتی شریف",
              degree: "کارشناسی ارشد",
              startDate: "1402",
              endDate: "1406",
            },
          ]
        }
      ]
    }
  },
}
</script>

<style scoped>

</style>
