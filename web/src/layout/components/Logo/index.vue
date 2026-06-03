<template>
  <div class="logo-wrap" :class="{ 'logo-collapsed': collapsed }">
    <img src="~@/assets/images/logo.png" alt="logo" class="logo-img" />
    <transition name="logo-text-fade">
      <div v-if="!collapsed" class="logo-text">
        <span class="logo-title">住一CRS</span>
        <span v-if="state.pmsmeun && propertyName" class="logo-sub" :title="propertyName">{{ propertyName }}</span>
      </div>
    </transition>
  </div>
</template>

<script lang="ts" setup>
import { reactive, ref, watch, onMounted, computed } from 'vue';
import { useRoute } from 'vue-router';
import { useUserStore } from '@/store/modules/user';
import { List } from '@/api/ten';
import { storage } from '@/utils/Storage';
import { getlang } from '@/utils/smjcomm';

export interface Props {
  collapsed?: boolean;
}
const props = withDefaults(defineProps<Props>(), { collapsed: false });

const userStore = useUserStore();
const currentRoute = useRoute();

const state = reactive({
  pmsmeun: false,
  chooseitem: userStore.getuserPms,
});

const propertyName = ref(
  userStore.getuserPms && userStore.getuserPms.id > 0
    ? namegetlang(userStore.getuserPms.nameLanguage)
    : ''
);

function namegetlang(data: any) {
  if (data) return getlang(data, userStore.language).content;
  return '';
}

watch(() => userStore.getuserPms, () => {
  propertyName.value = (userStore.getuserPms && userStore.getuserPms.id > 0)
    ? namegetlang(userStore.getuserPms.nameLanguage)
    : '';
}, { immediate: true, deep: true });

watch(() => currentRoute.fullPath, () => {
  state.pmsmeun = currentRoute.path.indexOf('pms/') !== -1;
});

const tenListLoad = async () => {
  const res = await List({ page: 1, pageSize: 200 });
  const wylists = res.list;
  storage.set('wylists', wylists);
  if (!userStore.getuserPms) {
    chooseItem(res.list[0]);
  }
};

const chooseItem = (item: any) => {
  state.chooseitem = item;
  storage.set('userPms', item);
  userStore.setuserPms(item);
};

onMounted(async () => {
  await tenListLoad();
  state.pmsmeun = currentRoute.path.indexOf('pms/') !== -1;
});
</script>

<style lang="less" scoped>
.logo-wrap {
  display: flex;
  align-items: center;
  height: 56px;
  padding: 0 16px;
  overflow: hidden;
  flex-shrink: 0;
  border-bottom: 1px solid rgba(255, 255, 255, 0.07);
  transition: padding 0.2s ease;

  &.logo-collapsed {
    padding: 0;
    justify-content: center;
  }
}

.logo-img {
  width: 28px;
  height: 28px;
  border-radius: 6px;
  flex-shrink: 0;
  object-fit: contain;
}

.logo-text {
  margin-left: 10px;
  min-width: 0;
  overflow: hidden;
}

.logo-title {
  display: block;
  font-size: 14px;
  font-weight: 700;
  color: #ffffff;
  letter-spacing: -0.02em;
  white-space: nowrap;
  line-height: 1.3;
}

.logo-sub {
  display: block;
  font-size: 11px;
  color: rgba(255, 255, 255, 0.45);
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
  max-width: 130px;
  margin-top: 1px;
}

// Fade transition for logo text
.logo-text-fade-enter-active,
.logo-text-fade-leave-active {
  transition: opacity 0.15s ease;
}
.logo-text-fade-enter-from,
.logo-text-fade-leave-to {
  opacity: 0;
}
</style>
