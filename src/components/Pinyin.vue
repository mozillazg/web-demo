<template>
    <textarea-view :value.sync="text" @notify="handleInput"></textarea-view>
    <textarea-view :value.sync="result" readonly=true></textarea-view>
</template>

<script>
'use strict'
import Textarea from './Textarea.vue'

export default {
  name: 'Pinyin',
  components: {
    'textarea-view': Textarea
  },
  data: () => {
    return {
      text: '你好',
      result: 'ni hao'
    }
  },
  methods: {
    handleInput: function (value) {
      let url = '/api/v1/pinyin'
      let data = {han: this.text}
      this.$http.post(url, data).then((response) => {
        let result = response.json()
        this.$set('result', result.pinyin)
      })
    }
  }
}
</script>

