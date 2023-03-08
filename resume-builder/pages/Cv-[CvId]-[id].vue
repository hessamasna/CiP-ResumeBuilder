<template>
  <!-- <div class="mb-2 bg-black"> -->
  <!-- <div class="mb-10 bg-black"> -->
  <!--      todo fix style-->
  <v-row class="mt-3 text-center px-5 mx-5" justify="center" align="start">
    <v-col cols="12" sm="2">
      <v-select
          variant="outlined"
          hide-details=true
          :items="fonts"
          v-model="data.font_family"
          item-value="value"
          item-text="title"
      ></v-select>
    </v-col>
    <v-col cols="12" sm="1">
      <v-btn @click="showColorPicker = !showColorPicker" variant="outlined" append-icon="mdi-chevron-down" height="56">
        انتخاب رنگ
      </v-btn>
      <v-dialog
          v-model="showColorPicker"
          width="auto"
      >
        <v-card>
          <v-card-text>
            <v-color-picker
                v-model="data.color"
                hide-sliders
                hide-inputs
                show-swatches
            ></v-color-picker>
          </v-card-text>
          <v-card-actions>
            <v-btn  color="green" variant="flat" block @click="showColorPicker = false">ذخیره رنگ</v-btn>
          </v-card-actions>
        </v-card>
      </v-dialog>
    </v-col>

    <v-col cols="12" sm="2">
      <v-text-field hide-details=true variant="outlined" v-model="data.font_size" type="number" label="Number" append-outer-icon="add"
                    @click:append-outer="increment"
                     @click:prepend="decrement"/>
    </v-col>


  </v-row>
  <v-row justify="center" align="start" class="mb-4">
    <v-col cols="12" sm="2">
      <v-btn color="green" variant="flat" block @click="saveCv()">
        ذخیره
      </v-btn>
    </v-col>
  </v-row>
  <!-- </div> -->
  <cv1 :loading="loading" :data="data" v-if="cvTemplateId == 1"></cv1>
  <cv2 :loading="loading" :data="data" v-else-if="cvTemplateId == 2"></cv2>
  <cv4 :loading="loading" :data="data" v-else-if="cvTemplateId == 4"></cv4>
  <!-- </div> -->
</template>

<script>
export default {
  name: "Cv-[CvId]-[id]",
  async created() {
    this.loading = true;
    this.cvTemplateId = this.$route.params.CvId
    this.id = this.$route.params.id

    let api = 'http://localhost:3000/cv/get/' + this.id;

    let res = await $fetch(api, {
      method: 'GET',
      headers: {
        'access_token': this.$store.state.status.access_token,
        'refresh_token': this.$store.state.status.refresh_token
      },
    }).then(res => {
      //todo
      this.loading = false;
      this.data = res.result;
    }).catch(error => {
      // this.data = {
      //   image: "https://tailone.tailwindtemplate.net/src/img/dummy/avatar1.png",
      //   font: 'vazir-FD',
      //   color: "#ffa500",
      //   fontSize: 15,
      //   cv_temp: 1,
      //   cv_name: "رزومه اول",
      //   personal: {
      //     name: "سیپ سی ویq",
      //     cvTitle: "توسعه دهنده",
      //     phone: "09121112233",
      //     age:22,
      //     email: "Cv@cipCv.com",
      //     website: "www.cipCv.com",
      //     address: "Tehran",
      //     social: [
      //       {
      //         title: "linkedin",
      //         path: "CipCV",
      //         url: "https://linkedin.com/cipC"
      //       },
      //       {
      //         title: "instagram",
      //         path: "CipCV",
      //         url: "https://instagram.com/CipCv"
      //       },
      //     ]
      //   },
      //   aboutMe: "لورم ایپسوم متن ساختگی با تولید سادگی نامفهوم از صنعت چاپ، و با استفاده از طراحان گرافیک است، چاپگرها و متون بلکه روزنامه و مجله در ستون و سطرآنچنان که لازم است، و برای شرایط فعلی تکنولوژی مورد نیاز، و کاربردهای متنوع با هدف بهبود ابزارهای کاربردی می باشد، کتابهای زیادی در شصت و سه درصد گذشته حال و آینده، شناخت فراوان جامعه و متخصصان را می طلبد، تا با نرم افزارها شناخت بیشتری را برای طراحان رایانه ای علی الخصوص طراحان خلاقی، و فرهنگ پیشرو در زبان فارسی ایجاد کرد، در این صورت می توان امید داشت که تمام و دشواری موجود در ارائه راهکارها، و شرایط سخت تایپ به پایان رسد و زمان مورد نیاز شامل حروفچینی دستاوردهای اصلی، و جوابگوی سوالات پیوسته اهل دنیای موجود طراحی اساسا مورد استفاده قرار گیرد.",
      //   workHistory: [
      //     {
      //       title: 'FrontEnd Developer',
      //       startDate: '2021',
      //       endDate: '2022',
      //       place: "Sharif",
      //       description: "Lorem ipsum dolor sit amet, consectetur adipisicing elit. Dolorem explicabo impedit magni reiciendis suscipit! Adipisci deleniti earum eveniet incidunt ipsa libero modi, molestiae nobis pariatur quod repellat, repudiandae rerum sit!"
      //     },
      //     {
      //       title: 'توسعه دهنده فرانت',
      //       startDate: '1399',
      //       endDate: '1400',
      //       place: "دانشگاه صنعتی شریف",
      //       description: " زمان مورد نیاز شامل حروفچینی دستاوردهای اصلی، و جوابگوی سوالات پیوسته اهل دنیای موجود طراحی اساسا مورد استفاده قرار گیرد."
      //     }
      //   ],
      //   skills: [
      //     {
      //       title: "Java",
      //       amount: 80,
      //     },
      //     {
      //       title: 'JS',
      //       amount: 85,
      //     },
      //     {
      //       title: 'Ruby',
      //       amount: 20,
      //     }
      //   ],
      //   educations: [
      //     {
      //       title: "مهندسی کامپیوتر",
      //       place: "صنعتی شریف",
      //       degree: "کارشناسی",
      //       startDate: "1398",
      //       endDate: "1402",
      //     },
      //     {
      //       title: "مهندسی کامپیوتر",
      //       place: "صنعتی شریف",
      //       degree: "کارشناسی ارشد",
      //       startDate: "1402",
      //       endDate: "1406",
      //     },
      //   ]
      // }
      this.loading = true;
      this.loading = false;
      console.log(error)
    })

    // this.loading = false;

  },
  data() {
    return {
      foo: 0,
      cvTemplateId: 1,
      id: 2,
      showColorPicker: false,
      loading: true,
      picker: null,
      fonts: [
        {
          title: 'ایران سنس',
          value: 'IRANSans'
        }, {
          title: 'وزیر',
          value: 'vazir-FD'
        }, {
          title: 'یکان',
          value: 'Yekan'
        },
      ],
      data: {}
    }
  },
  methods: {
    increment() {
      if (this.foo === 80) return;

      this.foo = parseInt(this.foo, 10) + 1
    },
    decrement() {
      console.log(this.foo)
      if (this.foo === 0) return;

      this.foo = parseInt(this.foo, 10) - 1
    },
    async saveCv() {
      let api = 'http://localhost:3000/cv/update  ';

      let res = await $fetch(api, {
        method: 'PUT',
        headers: {
          'access_token': this.$store.state.status.access_token,
          'refresh_token': this.$store.state.status.refresh_token
        },
        body: JSON.stringify(this.data)
      }).then(res => {
        // console.log(res)
      }).catch(error => {
        console.log(error)
      })
    },

  }
}
</script>

<style scoped>

</style>
