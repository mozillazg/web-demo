<template>
  <textarea class="form-control" v-el:input v-model="value" @change="notify"
    :cols="cols"
    :rows="rows"
    :name="name"
    :readonly="readonly"
    :required="required"
    :disabled="disabled"
    :maxlength="maxlength"
    :placeholder="placeholder"
  ></textarea>
</template>

<script>
  export default {
    // name: 'TextareaBox',
    props: {
      value: {
        type: String,
        required: true,
        twoWay: true,
        default: ''
      }
    },
    methods: {
      notify: function () {
        this.$dispatch('notify', this.value)
      }
    }
  }
</script>
