<template>
  <div class="logo">
    <img src="~@/assets/images/logo.png" alt="" :class="{ 'mr-2': !collapsed }" />
    <h2 v-if="!state.pmsmeun" v-show="!collapsed" class="title">住一CRS</h2>
    <h2 v-else v-show="!collapsed" class="title nohuanhang" style="max-width: 100px">
      <span v-if="state.chooseitem.id">
        <n-tooltip placement="bottom" trigger="hover">
          <template #trigger>
            <span> {{ propertyName }}</span>
          </template>
          <span> {{ propertyName }}</span>
        </n-tooltip>
      </span>
      <span v-else>所有物业</span>
    </h2>

    <!-- 切换物业 -->

    <n-popover trigger="click" :show-arrow="false" placement="bottom" v-if="false">
      <template #trigger>
        <div
          :class="{ 'ml-5': !collapsed }"
          v-show="!collapsed && state.pmsmeun"
          style="cursor: pointer"
        >
          <n-icon size="18" color="#9E9E9E">
            <ListSharp />
          </n-icon>
        </div>
      </template>
      <div class="flex-column" style="width: 220px; max-height: 500px">
        <!-- <div class="pb-5 pt-1">
          <div class="search flex-row">
            <div>
              <n-icon size="16" color="#8B8B8B">
                <Search />
              </n-icon>
            </div>
          </div>
        </div> -->

<!--        <div class="flex-row pb-5 pt-5" @click="chooseItem({ id: '', name: '所有业务' })">-->
<!--          <img src="~@/assets/images/wyall.png" class="wyicon" />-->
<!--          <div class="flex-item">-->
<!--            <div class="c333 f15 fw">所有物业</div>-->
<!--            <div class="c999">{{ wylists.length }}总共</div>-->
<!--          </div>-->
<!--        </div>-->
        <div class="flex-item lh-3">
          <div
            class="flex-row pb-5 cpus"
            v-for="(item, index) in wylists"
            :key="index"
            style="cursor: pointer"
            @click="chooseItem(item)"
          >
            <img :src="item.icon" class="wyicon" />
            <div class="flex-item">
              <div class="f15 fw nohuanhang">
                <n-tooltip placement="bottom" trigger="hover">
                  <template #trigger>
                    <span> {{ namegetlang(item.nameLanguage) }} </span>
                  </template>
                  <span> {{ namegetlang(item.nameLanguage) }} </span>
                </n-tooltip>
              </div>
              <div class="c999">{{ item.mode }}</div>
            </div>
            <div class="lh-3 mr-2">
              <n-icon size="18" color="#9E9E9E">
                <EllipsisHorizontalOutline />
              </n-icon>
            </div>
          </div>
        </div>
        <!-- <div class="flex-row">
          <div class="flex-item" style="line-height: 3; color: #053dc8"> 创建物业 </div>
          <div class="lh-3 ml-2" style="cursor: pointer">
            <n-icon size="18" color="#9E9E9E">
              <ChevronForwardSharp />
            </n-icon>
          </div>
        </div> -->
      </div>
    </n-popover>
  </div>
</template>

<script lang="ts" setup>
  import {
    ListSharp,
    EllipsisHorizontalOutline,
    Search,
    ChevronForwardSharp,
  } from '@vicons/ionicons5';
  import { storage } from '@/utils/Storage';

  import { List } from '@/api/ten';
  import { useRoute, useRouter } from 'vue-router';
  import { ref, onMounted, reactive, watch } from 'vue';
  import { getlang } from '@/utils/smjcomm';

  import { useUserStore } from '@/store/modules/user';
  const userStore = useUserStore();

  export interface Props {
    collapsed: boolean | true;
  }
  const namegetlang = (data) => {
    if (data) {
      return getlang(data, userStore.language).content;
    } else {
      return '暂无物业名';
    }
  };
  const props = withDefaults(defineProps<Props>(), {
    collapsed: true,
  });
  // 当前路由
  const currentRoute = useRoute();
  const router = useRouter();
  const state = reactive({
    pmsmeun: false,
    chooseitem: userStore.getuserPms,
  });
  // 物业名称
  const propertyName = ref(userStore.getuserPms&&userStore.getuserPms.id > 0 ? namegetlang(userStore.getuserPms.nameLanguage) : '全部物业');
  let wylists = ref([]);

  // 加载表格数据
  const tenListLoad = async () => {
    const res = await List({ page: 1, pageSize: 200 });
    wylists = res.list;
    storage.set('wylists', wylists);
    if(!userStore.getuserPms){
      chooseItem(res.list[0])
    }
    console.log(wylists);
  };
  // 跟随页面路由变化，切换菜单选中状态
  watch(
    () => currentRoute.fullPath,
    () => {
      if (currentRoute.path.indexOf('pms/') != -1) {
        state.pmsmeun = true;
        console.log('这是pms路由', currentRoute);
      } else {
        state.pmsmeun = false;
      }
    }
  );
  watch(
    () => userStore.getuserPms,
    () => {
      propertyName.value = (userStore.getuserPms&&userStore.getuserPms.id > 0) ? namegetlang(userStore.getuserPms.nameLanguage) : '全部物业';
    },
    {
      immediate: true,
      deep: true,
    }
  );
  const chooseItem = (item) => {
    state.chooseitem = item;

    storage.set('userPms', item);
    userStore.setuserPms(item);
    console.log(item);
  };

  onMounted(async () => {
    await tenListLoad();
    if (currentRoute.path.indexOf('pms/') != -1) {
      state.pmsmeun = true;
      console.log('这是pms路由', currentRoute);
    } else {
      state.pmsmeun = false;
    }
  });
</script>

<style lang="less" scoped>
  .search {
    background-color: #f7f7f7;
    border-radius: 5px;
    padding: 5px 10px;
    line-height: 2;
  }
  .wyicon {
    width: 34px;
    height: 34px;
    margin-right: 10px;
  }
  .logo {
    display: flex;
    align-items: center;
    justify-content: center;
    height: 64px;
    line-height: 64px;
    overflow: hidden;
    white-space: nowrap;

    img {
      width: auto;
      height: 30px;
      // border-radius: 20px;
      background-color: #053dc8;
    }

    .title {
      margin-bottom: 0px;
    }
  }
</style>
