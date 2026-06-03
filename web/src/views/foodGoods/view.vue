<template>
  <div>
    <n-drawer v-model:show="showModal" :width="dialogWidth">
      <n-drawer-content title="套餐管理详情" closable>
        <n-spin :show="loading" description="请稍候...">
          <n-descriptions label-placement="left" class="py-2" column="1">
            <n-descriptions-item>
              <template #label>
                餐厅ID
              </template>
              {{ formValue.restaurantId }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                套餐名称
              </template>
              {{ formValue.goodsName }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                促销语
              </template>
              <span v-html="formValue.introduction"></span>
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                标签（多选）
              </template>
              {{ formValue.labelIds }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                图集
              </template>
              <n-image-group>
                <n-space>
                  <span v-for="(item, key) in formValue?.images" :key="key">
                    <n-image style="margin-left: 10px; height: 100px; width: 100px" :src="item" />
                  </span>
                </n-space>
              </n-image-group>
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                套餐售价
              </template>
              {{ formValue.price }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                套餐原价
              </template>
              {{ formValue.marketPrice }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                单次最大可预定数
              </template>
              {{ formValue.maxOrderNum }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                用餐停留时间
              </template>
              {{ formValue.timeDuration }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                套餐详情
              </template>
              <span v-html="formValue.goodsContent"></span>
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                注意事项
              </template>
              <span v-html="formValue.notice"></span>
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                状态（1.正常2下架）
              </template>
              {{ formValue.goodsState }}
            </n-descriptions-item>
          </n-descriptions>
        </n-spin>
      </n-drawer-content>
    </n-drawer>
  </div>
</template>

<script lang="ts" setup>
import { computed, ref } from 'vue';
import { useMessage } from 'naive-ui';
import { View } from '@/api/foodGoods';
import { State, newState } from './model';
import { adaModalWidth } from '@/utils/hotgo';

const message = useMessage();
const loading = ref(false);
const showModal = ref(false);
const formValue = ref(newState(null));
const dialogWidth = computed(() => {
  return adaModalWidth(580);
});
const fileAvatarCSS = computed(() => {
  return {
    '--n-merged-size': `var(--n-avatar-size-override, 80px)`,
    '--n-font-size': `18px`,
  };
});

//下载
function download(url: string) {
  window.open(url);
}

function openModal(state: State) {
  showModal.value = true;
  loading.value = true;
  View({ id: state.id })
    .then((res) => {
      formValue.value = res;
    })
    .finally(() => {
      loading.value = false;
    });
}

defineExpose({
  openModal,
});
</script>

<style lang="less" scoped></style>


