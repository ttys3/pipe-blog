<template>
  <div class="console" id="particles">
    <div class="card login__content" ref="content">
      <form action="" class="login__form">
        <p class="first-login-tips" v-if="isFirstLogin">
          default username: <code>admin</code> <br>
          default password: <code>admin</code>
        </p>
        <label for="username" class="login__label">username
        <input type="text" class="login__input" name="username" id="username" v-model="username" v-on:keyup.enter="login"/>
        </label>

        <label for="password" class="login__label">password
        <input type="password" class="login__input" name="password" id="password" v-model="password" v-on:keyup.enter="login"/>
        </label>

        <v-btn class="btn--small btn--info" @click="login">Login</v-btn>
      </form>
    </div>
  </div>
</template>

<script>
  import 'particles.js'
  import { initParticlesJS } from '~/plugins/utils'

  export default {
    data () {
      return {
        isFirstLogin: false,
        password: "",
        username: "",
      }
    },
    head () {
      return {
        title: this.$t('welcome', this.$store.state.locale) + ' - NanoBlog',
      }
    },
    methods: {
      async doLogin () {
        const responseData = await this.axios.post(`/login`, {
          username: this.username.trim(),
          password: this.password.trim(),
        })
        return responseData
      },
      async login (event) {
        event.preventDefault()
        this.$store.commit('setSnackBar', {
          snackBar: true,
          snackMsg: this.$t('processing', this.$store.state.locale),
          snackModify: 'success',
        })
        try {
          const rs = await this.doLogin()
          // console.log('rs: %o', rs)
          if (rs.code === 0) {
            this.$store.commit('setSnackBar', {
              snackBar: true,
              snackMsg: rs.msg,
              snackModify: 'success',
            })
            // need full reload
            window.location.href = "/admin/"
          } else {
            this.$store.commit('setSnackBar', {
              snackBar: true,
              snackMsg: rs.msg,
            })
          }
        } catch (e) {
          console.log(e)
        }
      },
    },
    mounted () {
      // console.log(this.$store.state)
      if (this.$store.state.role > 0) {
        window.location.href = "/admin/"
        this.$router.push('/admin/')
        return
      }
      if (this.$store.state.isInit === false) {
        this.isFirstLogin = true
      }
      initParticlesJS('particles')
    },
  }
</script>

<style lang="sass">
  .ft__12
    font-size: 12px !important
  .login
    &__intro
      text-align: left
      width: 300px
      margin: 0 auto
      ul
        margin-bottom: 0 !important
    &__space
      height: 10px
    &__checkbox
      margin: 0 20px
      color: #999
      a
        text-decoration: underline
        color: #67757c
    &__form
        padding: 2em
    &__input
        border-width: 1px
        height: 2em
        width: 15vw
        display: inline-block
    &__label
        display: block
        margin: 2em 5em
  .first-login-tips
      border: #94cfe0 2px solid
      border-radius: 3%
      height: auto
      width: 220px
      margin: 0 auto
  .first-login-tips code
      color: #da3e3e

  @media (max-width: 470px)
    .login__checkbox
      line-height: 18px !important
</style>
