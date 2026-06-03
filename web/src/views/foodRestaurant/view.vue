<template>
  <div>
    <n-drawer v-model:show="showModal" :width="dialogWidth">
      <n-drawer-content title="餐厅管理详情" closable>
        <n-spin :show="loading" description="请稍候...">
          <n-descriptions label-placement="left" class="py-2" column="1">
            <n-descriptions-item>
              <template #label>
                名称
              </template>
              {{ formValue.name }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                菜系（多选）
              </template>
              {{ formValue.cuisineIds }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                标签（多选）
              </template>
              {{ formValue.labelIds }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                合作类型ID
              </template>
              {{ formValue.cooperateTypeId }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                LOGO
              </template>
              <n-image style="margin-left: 10px; height: 100px; width: 100px" :src="formValue.logo"/>
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
                简介
              </template>
              <span v-html="formValue.content"></span>
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                联系电话
              </template>
              {{ formValue.phone }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                营业时间
              </template>
              {{ formValue.openTime }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                地区省级ID
              </template>
              {{ formValue.areaPid }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                地区市级ID
              </template>
              {{ formValue.areaId }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                详细地址
              </template>
              {{ formValue.address }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                谷歌纬度
              </template>
              {{ formValue.ggLat }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                谷歌经度
              </template>
              {{ formValue.ggLng }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                纬度
              </template>
              {{ formValue.lat }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                经度
              </template>
              {{ formValue.lng }}
            </n-descriptions-item>
            <n-descriptions-item label="营业状态">
              <n-tag :type="getOptionTag(options.open_status, formValue?.openStatus)" size="small" class="min-left-space">
                {{ getOptionLabel(options.open_status, formValue?.openStatus) }}
              </n-tag>
            </n-descriptions-item>
            <n-descriptions-item label="1、启用 2、禁用">
              <n-tag :type="getOptionTag(options.sys_normal_disable, formValue?.status)" size="small" class="min-left-space">
                {{ getOptionLabel(options.sys_normal_disable, formValue?.status) }}
              </n-tag>
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                预约日期  1每天 2自定义
              </template>
              {{ formValue.orderTimeType }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                预约日期自定义周数
              </template>
              {{ formValue.orderTimeWeek }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                定休日周数
              </template>
              {{ formValue.restTimeWeek }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                定休日固定日期
              </template>
              {{ formValue.restTimeDate }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                预约确定模式 1手动确认 2自动确认
              </template>
              {{ formValue.orderConfirmType }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                预约确认自定义小时数
              </template>
              {{ formValue.orderConfirmHour }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                早市开始时间
              </template>
              {{ formValue.amOpenTime }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                早市结束时间
              </template>
              {{ formValue.amCloseTime }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                晚市开始时间
              </template>
              {{ formValue.pmOpenTime }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                晚市结束时间
              </template>
              {{ formValue.pmCloseTime }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                时间间隔
              </template>
              {{ formValue.timeDuration }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                时段最大容纳预定数量  0不限制
              </template>
              {{ formValue.timeDurationMax }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                至少提前预约天数  0当日可约
              </template>
              {{ formValue.advanceOrderDay }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                最长预约天数
              </template>
              {{ formValue.maxOrderDay }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                是否允许取消  1允许  2不允许
              </template>
              {{ formValue.cancelPolicyOpen }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                订单确认前取消费用为订单的%，0则免费
              </template>
              {{ formValue.beforeConfirmCancelRate }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                订单确认后距离到店多少天
              </template>
              {{ formValue.afterConfirmCancelDay }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                订单确认后距离到店天数前取消费用为订单的%，0则免费
              </template>
              {{ formValue.afterConfirmCancelRate1 }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                订单确认后距离到店天数后取消费用为订单的%，0则免费
              </template>
              {{ formValue.afterConfirmCancelRate2 }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                最大席位数
              </template>
              {{ formValue.maxSeat }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                是否允许抽烟  1允许  2不允许
              </template>
              {{ formValue.canSmoking }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                门店抽成类型 1跟随系统  2自定义
              </template>
              {{ formValue.settlementType }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                结算模式ID
              </template>
              {{ formValue.settlementId }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                门店结算比例
              </template>
              {{ formValue.settlementRate }}
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
import { View } from '@/api/foodRestaurant';
import { State, newState, options } from './model';
import { adaModalWidth, getOptionLabel, getOptionTag } from '@/utils/hotgo';
import { getFileExt } from '@/utils/urlUtils';

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


