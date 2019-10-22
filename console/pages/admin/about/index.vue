<template>
  <div>
    <div class="card">
      <div class="card__body">
        <div class="fn__flex admin__about">
          <div class="about__side">
            <img src="~assets/images/logo.png"/>
          </div>
          <div class="fn__flex-1 vditor-reset">
            <h2 class="fn__clear" v-if="isLatest">
            <span class="fn__left">
              {{ $t('about1', $store.state.locale) }}
              <a :href="download" target="_blank">{{ version }}</a>
            </span>
              <iframe class="about__github fn__left"
                      src="https://ghbtns.com/github-btn.html?user=ttys3&repo=nanoblog&type=star&count=true&size=large"
                      frameborder="0" scrolling="0" width="160px" height="30px"></iframe>
            </h2>
            <h2 class="fn__clear" v-else>
            <span class="fn__left">
               {{ $t('about2', $store.state.locale) }}
              <a class="ft__danger" :href="download" target="_blank">{{ version }}</a>
            </span>
              <iframe class="about__github fn__left"
                      src="https://ghbtns.com/github-btn.html?user=ttys3&repo=nanoblog&type=star&count=true&size=large"
                      frameborder="0" scrolling="0" width="160px" height="30px"></iframe>
            </h2>
            <p v-html="$t('about4', $store.state.locale)"></p>
            <p v-html="$t('about3', $store.state.locale)"></p>
            <p>
              <a href="https://github.com/b3log/pipe-themes" target="_blank" class="about__link btn btn--success">
                <v-icon>github</v-icon>
                {{ $t('theme', $store.state.locale) }}
              </a>
              <a href="https://hacpai.com/article/1513761942333" target="_blank" class="about__link btn btn--success">
                <v-icon>file-text</v-icon>
                {{ $t('doc', $store.state.locale) }}
              </a>
              <a href="https://hacpai.com/tag/pipe" target="_blank" class="about__link btn btn--success">
                <v-icon>hacpai-logo2</v-icon>
                {{ $t('community', $store.state.locale) }}
              </a>
              <a href="https://b3log.org" target="_blank" class="about__link btn btn--success">
                <v-icon>b3log-logo2</v-icon>
                B3log
              </a>
            </p>
            <br>
          </div>
        </div>
      </div>
    </div>
    <br>
    <div class="card">
      <div class="card__title">
        <h3>❤️ 关于NanoDM</h3>
      </div>
      <div class="card__body">
        <p>
          <a href="https://nanodm.net">NanoDM</a>最初是基于Phicomm N1 (Amlogic S905)硬件开发的一款MiniNAS系统,
          随着使用人数的增多, 和日益增长的需求, 发现原有的打包ROM刷固件固件的方式不能再满足需求, 因此NanoDM也正在向docker化的方向前进,
          目前维护的docker 容器有：
        </p>

        <a href="https://hub.docker.com/r/80x86/filebrowser" target="_blank">filebrowser enhanced</a>、
        <a href="https://hub.docker.com/r/80x86/qbittorrent" target="_blank">qBittorrent</a>、
        <a href="https://hub.docker.com/r/80x86/baidupcs">baidupcs web</a>、
        <a href="https://hub.docker.com/r/80x86/lychee">迷你相册程序Lychee</a>、
        <a href="https://hub.docker.com/r/80x86/olaindex">OneDrive目录显示程序OLAINDEX</a>、
        <a href="https://hub.docker.com/r/80x86/typecho">typecho 博客程序</a>、
        &nbsp;等一系列项目。
        <br/> <br/>
      </div>
      <ul class="list" v-if="list.length > 0">
        <li
          v-for="item in list"
          :key="item.oId"
          class="fn__flex">
          <a :href="`https://hacpai.com/member/${item.paymentUserName}`"
             :aria-label="item.paymentUserName"
             v-if="item.paymentUserName"
             target="_blank"
             class="avatar avatar--mid avatar--space pipe-tooltipped pipe-tooltipped--n"
             :style="`background-image: url(${item.paymentUserAvatar48})`"></a>
          <img v-else
               class="avatar avatar--mid avatar--space"
               src="~assets/images/logo.png"/>
          <div class="fn__flex-1">
            <div class="fn__flex">
              <div class="fn__flex-1">
                <a v-if="item.paymentUserName" target="_blank"
                   class="list__title" :href="`https://hacpai.com/member/${item.paymentUserName}`">
                  {{ item.paymentUserName }}
                </a>
                <span v-else class="list__title">匿名好心人</span>
                <span class="ft__green">{{ item.paymentAmount }} RMB</span>
              </div>
              <time class="list__meta">{{ item.paymentTimeStr }}</time>
            </div>
            <div class="vditor__reset" v-html="item.paymentMemo"></div>
          </div>
        </li>
      </ul>
    </div>
  </div>
</template>

<script>
  export default {
    data () {
      return {
        isLatest: true,
        download: 'https://github.com/ttys3/nanoblog/releases',
        version: '',
        list: [],
      }
    },
    head () {
      return {
        title: `${this.$t('about', this.$store.state.locale)} - ${this.$store.state.blogTitle}`,
      }
    },
    async mounted () {
      this.version = this.$store.state.version
      
      // const responseData = await this.axios.get('check-version')
      // if (responseData) {
      //   this.$set(this, 'isLatest', this.$store.state.version === responseData.version)
      //   this.$set(this, 'download', responseData.download)
      //   this.$set(this, 'version', responseData.version)
      // }

      // const responseListData = await this.axios.get('/hp/apis/sponsors?format=json')
      // if (responseListData) {
      //   this.$set(this, 'list', responseListData.payments)
      // }
    },
  }
</script>

<style lang="sass">
  @import '~assets/scss/_variables'

  .admin__about
    .vditor-reset
      padding-left: 50px

      h2
        font-size: 1.4em

    .about
      &__side
        margin-left: 30px
        width: 162px
        align-self: center

      &__github
        border: 0
        margin-left: 20px

      &__link
        margin: 10px 40px 10px 0

        .icon
          border-right: 1px dotted #fff
          padding-right: 10px
          margin: 0 10px 0 0
          height: 18px
          width: 18px

  @media (max-width: 768px)
    .admin__about
      display: block

      .about__github
        float: none
        margin: 10px 0 0 0

      .about__side
        text-align: center
        margin: 15px
        width: auto

      .vditor-reset h2
        font-size: 1.2em
</style>
