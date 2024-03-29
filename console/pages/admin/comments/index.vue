<template>
  <div class="card">
    <div class="card__body" v-show="!isBatch">
      <v-text-field
        v-if="list.length > 0 || isSearch"
        @keyup.enter="isSearch = true; getList()"
        class="fn__flex-1"
        :label="$t('enterSearch', $store.state.locale)"
        v-model="keyword">
      </v-text-field>
      <div v-if="list.length < 1 && !isSearch">
        {{ $t('noData', $store.state.locale) }}
      </div>
    </div>
    <div class="card__batch-action fn__flex" v-show="isBatch">
      <label class="checkbox fn__flex-1">
        <input
          type="checkbox"
          :checked="isSelectAll"
          @click="selectAll"/>
        <span class="checkbox__icon"></span>
        {{ $t('hasChoose', $store.state.locale) }}
        {{selectedIds.length}}
        {{ $t('cntContent', $store.state.locale) }}
      </label>
      <div>
        <v-btn
          :class="{'btn--disabled': selectedIds.length === 0}"
          class="btn--danger"
          @click="batchAction">
          {{ $t('delete', $store.state.locale) }}
        </v-btn>
        <v-btn class="btn--success btn--space" @click="isBatch = false; selectedIds = []">
          {{ $t('cancel', $store.state.locale) }}
        </v-btn>
      </div>
    </div>
    <ul class="list" v-if="list.length > 0">
      <li
        :class="{'selected': isSelected(item.id)}"
        @click="setSelectedId(item.id)"
        v-for="item in list" :key="item.id" class="fn__flex">
        <a class="avatar avatar--mid avatar--space pipe-tooltipped pipe-tooltipped--s"
           :aria-label="item.author.name"
           :href="item.author.url"
           :style="`background-image: url(${item.author.avatarURL})`"></a>
        <div class="fn__flex-1">
          <div class="fn__flex">
            <div class="fn__flex-1">
              <a :href="item.url"
                 class="list__title">
                {{ item.title }}
              </a>
              <small class="fn-nowrap" v-if="userCount > 1">
                By <a :href="item.articleAuthor.url">{{ item.articleAuthor.name }}</a>
              </small>
            </div>
            <div>
              <v-btn
                v-show="!isBatch"
                v-if="$store.state.name === item.author.name || $store.state.role < 3"
                class="btn btn--danger btn--small"
                @click.stop="remove(item.id)">{{ $t('delete', $store.state.locale) }}
              </v-btn>
            </div>
          </div>
          <div class="vditor-reset" v-html="item.content"></div>
          <div class="list__meta">
            <time class="fn-nowrap">{{ item.createdAt }}</time>
          </div>
        </div>
      </li>
    </ul>
    <ul class="list" v-else-if="isSearch && !isBatch">
      <li class="fn__flex fn__pointer" @click="isSearch = false;keyword = '';getList();">
        {{ $t('searchNull', $store.state.locale) }} &nbsp;
        <span class="ft__danger">{{keyword}}</span>
      </li>
    </ul>
    <div class="pagination--wrapper fn__clear" v-if="pageCount > 1">
      <v-pagination
        :length="pageCount"
        v-model="currentPageNum"
        :total-visible="windowSize"
        class="fn__right"
        circle
        next-icon="angle-right"
        prev-icon="angle-left"
        @input="getList"
      ></v-pagination>
    </div>
  </div>
</template>

<script>
  import Vue from 'vue'
  import {LazyLoadImage} from '~/plugins/utils'

  export default {
    data () {
      return {
        isSearch: false,
        isSelectAll: false,
        isBatch: false,
        selectedIds: [],
        currentPageNum: 1,
        pageCount: 1,
        windowSize: 1,
        list: [],
        userCount: 1,
        keyword: ''
      }
    },
    head () {
      return {
        title: `${this.$t('commentList', this.$store.state.locale)} - ${this.$store.state.blogTitle}`
      }
    },
    methods: {
      selectAll () {
        this.$set(this, 'isSelectAll', !this.isSelectAll)
        if (!this.isSelectAll) {
          this.$set(this, 'selectedIds', [])
          return
        }

        const selectedIds = []
        this.list.forEach((data) => {
          selectedIds.push(data.id)
        })
        this.$set(this, 'selectedIds', selectedIds)
      },
      isSelected (id) {
        let isSelected = false
        this.selectedIds.forEach((data) => {
          if (data === id) {
            isSelected = true
          }
        })
        return isSelected
      },
      async batchAction () {
        if (this.selectedIds.length === 0) {
          return
        }
        if (!confirm(this.$t('confirmDelete', this.$store.state.locale))) {
          return
        }
        const responseData = await this.axios.post('/console/comments/batch-delete', {
          ids: this.selectedIds
        })
        if (responseData.code === 0) {
          this.$set(this, 'error', false)
          this.$set(this, 'errorMsg', '')
          this.$set(this, 'isSelectAll', false)
          this.$set(this, 'selectedIds', [])
          this.getList()
        } else {
          this.$set(this, 'error', true)
          this.$set(this, 'errorMsg', responseData.msg)
        }
      },
      setSelectedId (id) {
        let isSelected = false
        this.selectedIds.forEach((data) => {
          if (data === id) {
            isSelected = true
          }
        })

        if (isSelected) {
          this.$set(this, 'selectedIds', this.selectedIds.filter((data) => id !== data))
          if (this.selectedIds.length < 1) {
            this.$set(this, 'isBatch', false)
          }
        } else {
          this.$set(this, 'isBatch', true)
          this.selectedIds.push(id)
        }

        if (this.selectedIds.length === this.list.length) {
          this.$set(this, 'isSelectAll', true)
        } else {
          this.$set(this, 'isSelectAll', false)
        }
      },
      async getList (currentPage = 1) {
        const responseData = await this.axios.get(`/console/comments?p=${currentPage}&key=${this.keyword}`)
        if (responseData) {
          this.$set(this, 'userCount', responseData.userCount)
          this.$set(this, 'list', responseData.comments || [])
          this.$set(this, 'currentPageNum', responseData.pagination.currentPageNum)
          this.$set(this, 'pageCount', responseData.pagination.pageCount)
          this.$set(this, 'windowSize', document.documentElement.clientWidth < 721 ? 5 : responseData.pagination.windowSize)
          Vue.nextTick(() => {
            LazyLoadImage()
          })
        }
      },
      async remove (id) {
        if (!confirm(this.$t('confirmDelete', this.$store.state.locale))) {
          return
        }
        const responseData = await this.axios.delete(`/console/comments/${id}`)
        if (responseData === null) {
          this.$store.commit('setSnackBar', {
            snackBar: true,
            snackMsg: this.$t('deleteSuccess', this.$store.state.locale),
            snackModify: 'success'
          })
          this.getList()
        }
      }
    },
    mounted () {
      this.getList()
    }
  }
</script>
