<template>
  <section>
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
          <p class="help">グループ名はDoorkeeperのグループページのURLの https://<b>XXXXX</b>.doorkeeper.jp/ の太字の部分を入力してください</p>
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
            <button class="button is-info" @click="copyToClipboard">
              <i class="far fa-clipboard"></i>
              Copy
            </button>
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
  </section>
</template>

<script>
export default {
  data() {
    return {
      site: "doorkeeper",
      format: "atom",
      groupName: "trbmeetup",
    };
  },
  computed: {
    feedUrl() {
      return window.location.protocol + "//" + window.location.host + "/api/" + this.site + "/" + this.groupName.trim() + "." + this.format;
    },
  },
  methods: {
    copyToClipboard() {
      const feedUrlInput = document.getElementById('feedUrl');
      feedUrlInput.select();
      document.execCommand('copy');
    },
  },
};
</script>
