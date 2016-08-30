<template>
    <textarea-view :value.sync="pinyin" @notify="handleInput"></textarea-view>
    <textarea-view :value.sync="result" readonly=true></textarea-view>
</template>

<script>
import Textarea from './Textarea.vue'

export default {
  name: 'Stpinyin',
  components: {
    'textarea-view': Textarea
  },
  data: function () {
    return {
      pinyin: 'pinyin',
      result: ''
    }
  },
  methods: {
    handleInput: function (value) {
      let url = '/api/v1/stpinyin'
      let data = {pinyin: this.pinyin}
      this.$http.post(url, data).then((response) => {
        let result = response.json()
        this.$set('result', result.pinyin)
      })
    }
  }
}
</script>
