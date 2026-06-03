<template>
  <div class="flex-row">
    <!-- <div class="bgfff text-c" style="width: 140px; border-radius: 5px; margin-right: 10px"> -->
      <!-- <n-menu
        :options="state.menuOptions"
        :value="state.valuekey"
        :default-expand-all="true"
        @update:value="onmenuOptions"
      /> -->
      <!-- <img style="width: 110px; margin: 15px auto" src="@/assets/images/ltk.png" /> -->
      <!-- <div
        :class="state.valuekey == item.key ? 'daohang1click' : 'daohang1'"
        v-for="(item, index) in state.menuOptions"
        :key="index"
        @click="onmenuOptions(item.key)"
      >
        <div class="lan" v-if="state.valuekey == item.key"></div>
        <n-icon
          size="18"
          :color="state.valuekey == item.key ? '#053dc8ff' : '#333'"
          style="margin-right: 5px"
        >
          <component :is="item.icon" />
        </n-icon>
        {{ item.label }}
      </div> -->
    <!-- </div> -->

    <iframe :src="state.iframeurl" class="flex-item" style="height: 88vh" />
  </div>
</template>

<script lang="ts" setup>
  import { h, reactive, ref, computed, onMounted } from 'vue';
  import { SpeedometerOutline, LogoWechat } from '@vicons/ionicons5';
  import { kefuoss } from '@/api/pmsProperty';
  import { useRouter } from 'vue-router';

  const router = useRouter();
  const state = reactive({
    iframeurl: '',
    valuekey: 'chat_main',
    menuOptions: [
      // {
      //   label: '面板',
      //   key: 'mainGuide',
      //   icon: SpeedometerOutline,
      // },
      {
        label: '消息',
        key: 'chat_main',
        icon: LogoWechat,
      },
    ],
  });
  const onmenuOptions = (key: string) => {
    state.valuekey = key;
    LoadUrl();
  };
  const LoadUrl = () => {
    kefuoss({ redirect: '/' + state.valuekey }).then((res) => {
      state.iframeurl = res.iframeUrl;
    });
  };
  onMounted(() => {
    if (router.currentRoute.value.query?.key) {
      state.valuekey = router.currentRoute.value.query.key;
    }
    LoadUrl();
  });
</script>

<style lang="less" scoped>
  .daohang1 {
    cursor: pointer;
    position: relative;
    width: 90%;
    padding: 10px 0px 10px 20px;
    margin-bottom: 10px;
    color: #333;
    text-align: left;
  }
  .daohang1click {
    cursor: pointer;
    position: relative;
    width: 90%;
    padding: 10px 0px 10px 20px;
    margin-bottom: 10px;
    text-align: left;
    background-color: #053dc814;
    color: #053dc8ff;
    border-top-right-radius: 15px;
    border-bottom-right-radius: 15px;
    font-weight: 550;
  }
  .daohang1:hover {
    background-color: #053dc814;
    color: #053dc8ff;
    border-top-right-radius: 15px;
    border-bottom-right-radius: 15px;
  }
  .lan {
    position: absolute;
    left: 0px;
    top: 0px;
    width: 4px;
    height: 100%;
    background-color: #053dc8ff;
  }
</style>
