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
              <input class="input" type="text" v-model="groupName">
            </div>
            <p class="help">グループ名はconnpassだとグループページのURLの https://<b>XXXXX</b>.connpass.com/ の太字の部分、DoorkeeperだとグループページのURLの https://<b>XXXXX</b>.doorkeeper.jp/ の太字の部分を入力してください</p>
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
            <div class="control">
              <i v-if="loadingStatus == 'loading'" class="fa fa-2x fa-spinner fa-spin"></i>
              <i v-else-if="loadingStatus == 'success'" class="fa fa-2x fa-thumbs-up"></i>
              <i v-else-if="loadingStatus == 'error'" class="fa fa-2x fa-exclamation-triangle"></i>
            </div>
          </div>
        </div>
      </div>

      <div class="field is-horizontal">
        <div class="field-label">
        </div>
        <div class="field-body">
          <div class="field">
            <p class="help">このURLをRSSリーダーやGoogleカレンダーなどに登録してください</p>
          </div>
        </div>
      </div>
    </div>
  </section>
</template>

<script lang="ts">
  import { Vue, Component, Watch } from 'vue-property-decorator'
  import axios from 'axios';
  import lodash from 'lodash';

  @Component
  export default class App extends Vue {
    site = "connpass";
    format = "atom";
    groupName = "gocon";
    loadingStatus = "";
    debouncedCheckFeed = lodash.debounce(this.checkFeed, 500);

    get feedUrl() {
      return window.location.protocol + "//" + window.location.host + "/api/" + this.site + "/" + this.groupName.trim() + "." + this.format
    }

    @Watch("feedUrl")
    onFeedUrlChanged(newFeedUrl: string, oldFeedUrl: string) {
      if (newFeedUrl == oldFeedUrl) {
        return
      }
      this.loadingStatus = "loading";
      this.debouncedCheckFeed();
    }

    async checkFeed() {
      try {
        const res = await axios.head(this.feedUrl);
        if(res.status == 200){
          this.loadingStatus = "success";
        } else {
          this.loadingStatus = "error";
        }
      } catch (error) {
        this.loadingStatus = "error";
      }
    }
  }
</script>
