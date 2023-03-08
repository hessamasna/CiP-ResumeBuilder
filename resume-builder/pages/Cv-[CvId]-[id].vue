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
      <v-text-field hide-details=true variant="outlined" v-model.number="data.font_size" type="number" label="Number" append-outer-icon="add"
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
      snackbar: {
        color: null,
        show: false,
        message: "با موفقیت پاک شد"
      },
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
        this.snackbar = {
          color: 'green',
          show: true,
          message: 'با موفقیت ذخیره شد'
        }
        // console.log(res)
      }).catch(error => {
        this.snackbar = {
          color: 'red',
          show: true,
          message: 'خطایی رخ داده است'
        }
        console.log(error)
      })
    },

  }
}
</script>

<style scoped>

</style>
