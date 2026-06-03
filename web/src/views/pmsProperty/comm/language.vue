<template>
  <div>
    <n-tabs type="line" :default-value="props.tabstype" animated>
      <n-tab-pane tab="日本语" name="ja">
        <div
          class="p-1 pl-10 lh-3 nohuanghang"
          style="border: 1px solid #e7e7e7; height: 80px; overflow-y: auto"
        >
     
          <span v-if="state.langdatadata['ja'].listok">
            <a-tag
              color="blue"
              class="mr-2"
              v-for="item in state.langdatadata['ja'].content"
              closable
              >{{ item }}
            </a-tag>
          </span>
          <span v-else>
            <div v-if="state.langdatadata['ja'].content">
              {{ state.langdatadata['ja'].content }}
            </div>
            <div v-else>言語の配置はまだありません</div>
          </span>
        </div>
      </n-tab-pane>
      <n-tab-pane tab="English" name="en">
        <div
          class="p-1 pl-10 lh-3 nohuanghang"
          style="border: 1px solid #e7e7e7; height: 80px; overflow-y: auto"
        >
          <span v-if="state.langdatadata['en'].listok">
            <a-tag
              color="blue"
              class="mr-2"
              v-for="item in state.langdatadata['en'].content"
              closable
              >{{ item }}</a-tag
            >
          </span>
          <span v-else>
            <div v-if="state.langdatadata['en'].content">
              {{ state.langdatadata['en'].content }}
            </div>
            <div v-else>No language configuration currently available</div>
          </span>
        </div>
      </n-tab-pane>
      <n-tab-pane tab="简体中文" name="zh">
        <div
          class="p-1 pl-10 lh-3 nohuanghang"
          style="border: 1px solid #e7e7e7; height: 80px; overflow-y: auto"
        >
          <span v-if="state.langdatadata['zh'].listok">
            <a-tag
              color="blue"
              class="mr-2"
              v-for="item in state.langdatadata['zh'].content"
              closable
              >{{ item }}</a-tag
            >
          </span>
          <span v-else>
            <div v-if="state.langdatadata['zh'].content">
              {{ state.langdatadata['zh'].content }}
            </div>
            <div v-else>暂无语言配置</div>
          </span>
        </div>
      </n-tab-pane>
      <n-tab-pane tab="한국어" name="ko">
        <div
          class="p-1 pl-10 lh-3 nohuanghang"
          style="border: 1px solid #e7e7e7; height: 80px; overflow-y: auto"
        >
          <span v-if="state.langdatadata['ko'].listok">
            <a-tag
              color="blue"
              class="mr-2"
              v-for="item in state.langdatadata['ko'].content"
              closable
              >{{ item }}</a-tag
            >
          </span>
          <span v-else>
            <div v-if="state.langdatadata['ko'].content">
              {{ state.langdatadata['ko'].content }}
            </div>
            <div v-else>언어 설정이 없습니다</div>
          </span>
        </div>
      </n-tab-pane>
      <n-tab-pane tab="繁体中文" name="zh_CN">
        <div
          class="p-1 pl-10 lh-3 nohuanghang"
          style="border: 1px solid #e7e7e7; height: 80px; overflow-y: auto"
        >
          <span v-if="state.langdatadata['zh_CN']&&state.langdatadata['zh_CN'].listok">
            <a-tag
              color="blue"
              class="mr-2"
              v-for="item in state.langdatadata['zh_CN'].content"
              closable
              >{{ item }}</a-tag
            >
          </span>
          <span v-else>
            <div v-if="state.langdatadata['zh_CN']&&state.langdatadata['zh_CN'].content">
              {{ state.langdatadata['zh_CN'].content }}
            </div>
            <div v-else>暫無語言配置</div>
          </span>
        </div>
      </n-tab-pane>
    </n-tabs>
  </div>
</template>

<script lang="ts" setup>
  import { reactive, onMounted } from 'vue';
  const state = reactive({
    langdatadata: {
      ja: { listok: false },
      en: { listok: false },
      zh: { listok: false },
      ko: { listok: false },
      zh_CN: { listok: false },
    },
  });

  // 对象写法并校验
  const props = defineProps({
    tabstype: String,
    langdata: Object,
    height: String,
  });
  onMounted(() => {
    if (props.langdata.zh) {      
      state.langdatadata = props.langdata;
    }
  });
</script>

<style lang="less" scoped></style>
