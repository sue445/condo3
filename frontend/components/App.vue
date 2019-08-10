<template>
  <section>
    <div class="container">
      <div class="field is-horizontal">
        <div class="field-label">
          <label class="label">サイト</label>
        </div>
        <div class="field-body">
          <div class="field is-grouped is-narrow">
            <div class="control">
              <label class="radio">
                <input type="radio" name="site" value="connpass" v-model="site" checked>
                connpass
              </label>
              <label class="radio">
                <input type="radio" name="site" value="doorkeeper" v-model="site">
                Doorkeeper
              </label>
            </div>
          </div>
        </div>
      </div>

      <div class="field is-horizontal">
        <div class="field-label">
          <label class="label">フォーマット</label>
        </div>
        <div class="field-body">
          <div class="field is-grouped is-narrow">
            <div class="control">
              <label class="radio">
                <input type="radio" name="format" value="atom" v-model="format" checked>
                Atom
              </label>
              <label class="radio">
                <input type="radio" name="format" value="ics" v-model="format">
                iCalendar
              </label>
            </div>
          </div>
        </div>
      </div>

      <div class="field is-horizontal">
        <div class="field-label">
          <label class="label">グループ名</label>
        </div>
        <div class="field-body">
          <div class="field">
            <div class="control">
              <input class="input is-danger" type="text" v-model="groupName">
            </div>
          </div>
        </div>
      </div>

      <div class="field is-horizontal">
        <div class="field-label">
          <label class="label">URL</label>
        </div>
        <div class="field-body">
          <div class="field has-addons">
            <div class="control is-expanded">
              <input class="input" type="text" :value="feedUrl" id="feedUrl" readonly>
            </div>
            <div class="control">
              <a class="button is-info" data-clipboard-target="#feedUrl">
                <i class="far fa-clipboard"></i>
                Copy
              </a>
            </div>
          </div>
        </div>
      </div>
    </div>
  </section>
</template>

<script lang="ts">
  import Vue from 'vue'
  import Component from 'vue-class-component'

  @Component
  export default class App extends Vue {
    site = "connpass";
    format = "atom";
    groupName = "gocon";

    get feedUrl() {
      return 'https://condo3.appspot.com/api/' + this.site + "/" + this.groupName + "." + this.format
    }
  }
</script>
