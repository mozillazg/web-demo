<template>
    <textarea-view :value.sync="text" @notify="handleInput"></textarea-view>
    <textarea-view :value.sync="result" readonly=true></textarea-view>
</template>

<script>
import Textarea from './Textarea.vue'

export default {
  name: 'Slugify',
  components: {
    'textarea-view': Textarea
  },
  data: function () {
    return {
      text: 'text',
      result: 'text'
    }
  },
  methods: {
    handleInput: function (value) {
      let url = '/api/v1/slugify'
      let data = {'string': this.text}
      this.$http.post(url, data).then((response) => {
        let result = response.json()
        this.$set('result', result.slug)
      })
    }
  }
}
</script>

