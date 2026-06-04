<template>
  <div class="member-admin-page">
    <n-spin :show="loading" description="请稍候...">
      <n-card
        :bordered="false"
        class="proCard"
        size="small"
        :segmented="{ content: true }"
      >
        <n-form
          ref="formRef"
          :model="formValue"
          :rules="rules"
          :label-placement="settingStore.isMobile ? 'top' : 'left'"
          :label-width="200"
          class="py-4"
        >
          <n-grid cols="1 s:1 m:1 l:1 xl:1 2xl:1" responsive="screen" y-gap="10">
            <n-gi span="1">
              <n-form-item label="是否开启下单奖励积分" path="isGetOpen">
                <n-switch unchecked-value="N" checked-value="Y" v-model:value="formValue.isGetOpen" @update:value="isGetOpenChange"/>
                <template #feedback v-if="parseInt(sceneId) == 1">开启后，用户下单预定民宿，在退房完成后将获得积分奖励</template>
                <template #feedback v-else-if="parseInt(sceneId) == 2">开启后，用户下单预定套餐，在核销完成后将获得积分奖励</template>
                <template #feedback v-else-if="parseInt(sceneId) == 3">开启后，用户下单预定按摩服务，在订单完成后将获得积分奖励</template>
                <template #feedback v-else-if="parseInt(sceneId) == 4">开启后，用户下单预定接送机/包车服务，在订单完成后将获得积分奖励</template>
              </n-form-item>
            </n-gi>
            <n-gi span="1" v-if="isGetOpenFlag">
              <n-form-item label="下单奖励积分比例" path="getRate">
                <n-input-number placeholder="建议输入正整数" v-model:value="formValue.getRate" :show-button="false" :precision="0">
                  <template #suffix>
                    %
                  </template>
                </n-input-number>
                <template #feedback v-if="parseInt(sceneId) == 1">比率必须为0-100的整数，例：当设置为1时，用户民宿订单每消费100JPY奖励1个积分</template>
                <template #feedback v-else-if="parseInt(sceneId) == 2">比率必须为0-100的整数，例：当设置为1时，用户餐饮订单每消费100JPY奖励1个积分</template>
                <template #feedback v-else-if="parseInt(sceneId) == 3">比率必须为0-100的整数，例：当设置为1时，用户按摩订单每消费100JPY奖励1个积分</template>
                <template #feedback v-else-if="parseInt(sceneId) == 4">比率必须为0-100的整数，例：当设置为1时，用户接送机/包车订单每消费100JPY奖励1个积分</template>
                <template #feedback v-else-if="parseInt(sceneId) == 5">比率必须为0-100的整数，例：当设置为1时，用户储物柜订单每消费100JPY奖励1个积分</template>
              </n-form-item>
            </n-gi>
            <n-gi span="1">
              <n-form-item label="下单是否允许积分抵扣" path="isPayOpen">
                <n-switch unchecked-value="N" checked-value="Y" v-model:value="formValue.isPayOpen" @update:value="isPayOpenFlagChange"/>
                <template #feedback v-if="parseInt(sceneId) == 1">开启后，用户下单预定民宿，可以使用积分抵扣一部分订单金额</template>
                <template #feedback v-else-if="parseInt(sceneId) == 2">开启后，用户下单预定套餐，可以使用积分抵扣一部分订单金额</template>
                <template #feedback v-else-if="parseInt(sceneId) == 3">开启后，用户下单预定按摩服务，可以使用积分抵扣一部分订单金额</template>
                <template #feedback v-else-if="parseInt(sceneId) == 4">开启后，用户下单预定接送机/包车服务，可以使用积分抵扣一部分订单金额</template>
                <template #feedback v-else-if="parseInt(sceneId) == 5">开启后，用户下单预定储物柜服务，可以使用积分抵扣一部分订单金额</template>
              </n-form-item>
            </n-gi>
            <n-gi span="1" v-if="isPayOpenFlag">
              <n-form-item label="订单金额门槛" path="limitMoney">

                <n-radio-group v-model:value="formValue.isLimitMoney" name="isLimitMoney"  @update:value="isLimitMoneyChange">
                  <n-space>
                    <n-radio :value="1">不限制</n-radio>
                    <n-radio :value="2">限制</n-radio>
                  </n-space>
                </n-radio-group>
              </n-form-item>
            </n-gi>

            <n-gi span="1" v-if="isPayOpenFlag&&isLimitMoneyFlag">
              <n-form-item label="订单金额超过" path="limitMoney">
                <n-input-number placeholder="建议输入正整数" v-model:value="formValue.limitMoney" :show-button="false"/> <span style="padding-left: 5px">JPY，才可以使用积分抵现</span>
              </n-form-item>
            </n-gi>

            <n-gi span="1" v-if="isPayOpenFlag">
              <n-form-item label="积分抵现上限" path="isPayRate">
                <n-radio-group v-model:value="formValue.isPayRate" name="isPayRate" @update:value="isPayRateFlagChange">
                  <n-space>
                    <n-radio :value="1">不限制</n-radio>
                    <n-radio :value="2">限制</n-radio>
                  </n-space>
                </n-radio-group>
              </n-form-item>
            </n-gi>

            <n-gi span="1" v-if="isPayOpenFlag&&isPayRateFlag">
              <n-form-item label="每笔订单最多抵扣订单金额的" path="payRate">
                <n-input-number placeholder="建议输入正整数" v-model:value="formValue.payRate" :show-button="false" >
                  <template #suffix>
                    %
                  </template>
                </n-input-number>
              </n-form-item>
            </n-gi>
          </n-grid>
          <div style="text-align: center">
            <n-space justify="center">
              <n-button type="info" :loading="formBtnLoading" @click="confirmForm">
                确定
              </n-button>
            </n-space>
          </div>
        </n-form>
      </n-card>
    </n-spin>
  </div>
</template>

<script lang="ts" setup>
import {ref, computed, onMounted} from 'vue';
import { Edit, View } from '@/api/pmsMemberScene';
import { options, State, newState, rules } from './model';
import { useProjectSettingStore } from '@/store/modules/projectSetting';
import { useMessage } from 'naive-ui';
import { adaModalWidth } from '@/utils/hotgo';

interface Props {
  sceneId?: string;
}

const props = withDefaults(defineProps<Props>(), {
  sceneId: '0',
});
const emit = defineEmits(['reloadTable']);
const message = useMessage();
const settingStore = useProjectSettingStore();
const loading = ref(false);
const showModal = ref(false);
const formValue = ref<State>(newState(null));
const formRef = ref<any>({});
const formBtnLoading = ref(false);
const dialogWidth = computed(() => {
  return adaModalWidth(640);
});
const isLimitMoneyFlag = ref(false);
const isPayRateFlag = ref(false)
const isGetOpenFlag = ref(false)
const isPayOpenFlag = ref(false)

function confirmForm(e) {
  e.preventDefault();
  formBtnLoading.value = true;
  formRef.value.validate((errors) => {
    if (!errors) {

      if(formValue.value.isLimitMoney == 1){
        formValue.value.limitMoney = 0
      }

      if(formValue.value.isPayRate == 1){
        // 不限制积分抵扣上限，则积分抵扣比例为100%
        formValue.value.payRate = 100
      }

      Edit(formValue.value).then((_res) => {
        message.success('操作成功');
        setTimeout(() => {

          emit('reloadTable');
        });
      });
    } else {
      message.error('请填写完整信息');
    }
    formBtnLoading.value = false;
  });
}

function isGetOpenChange(value) {
  console.log('value', value)
  if(value == 'Y'){
    isGetOpenFlag.value = true
  }else{
    isGetOpenFlag.value = false
  }
}

function isPayOpenFlagChange(value) {
  console.log('value', value)
  if(value == 'Y'){
    isPayOpenFlag.value = true
  }else{
    isPayOpenFlag.value = false
  }
}

function isLimitMoneyChange(value) {
  console.log('value', value)
  if(value == 1){
    isLimitMoneyFlag.value = false
  }else{
    isLimitMoneyFlag.value = true
  }
}

function isPayRateFlagChange(value) {
  console.log('value', value)
  if(value == 1){
    isPayRateFlag.value = false
  }else{
    isPayRateFlag.value = true
  }
}


onMounted(() => {
  load();
});

function load() {
  // 编辑
  loading.value = true;
  View({ id: props.sceneId })
    .then((res) => {

      res.isLimitMoney = 1
      isLimitMoneyFlag.value = false
      if (res.limitMoney > 0){
        res.isLimitMoney = 2
        isLimitMoneyFlag.value = true
      }
      res.isPayRate = 1
      isPayRateFlag.value = false
      if (res.payRate > 0){
        isPayRateFlag.value = true
        res.isPayRate = 2
      }
      if (res.isGetOpen == 'Y'){
        isGetOpenFlag.value = true
      }else{
        isGetOpenFlag.value = false
      }

      if (res.isPayOpen == 'Y'){
        isPayOpenFlag.value = true
      }else{
        isPayOpenFlag.value = false
      }

      formValue.value = res;
      console.log('----------------------------',formValue.value)
    })
    .finally(() => {
      loading.value = false;
    });
}

</script>

<style lang="less"></style>

