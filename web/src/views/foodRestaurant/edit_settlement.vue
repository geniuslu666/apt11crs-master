<template>
  <div>
    <n-spin :show="show" description="请稍候...">
      <n-form
        ref="formRef"
        :model="formValue"
        :rules="rules"
        :label-placement="settingStore.isMobile ? 'top' : 'left'"
        :label-width="150"
        class="py-4"
      >
        <n-grid cols="1 s:1 m:1 l:1 xl:1 2xl:1" responsive="screen">
          <n-gi span="1">
            <n-form-item label="结算模式" path="settlementId">
              <n-select
                v-model:value="formValue.settlementId"
                :options="settlementList"
                label-field="name"
                value-field="id"
                style="width: 200px"
              />
            </n-form-item>
          </n-gi>
          <n-gi span="1">
            <n-form-item label="门店抽成" path="settlementType">
              <n-radio-group v-model:value="formValue.settlementType" name="settlementType">
                <n-space>
                  <n-radio :value="1">
                    跟随系统
                  </n-radio>
                  <n-radio :value="2">
                    自定义
                  </n-radio>
                </n-space>
              </n-radio-group>
            </n-form-item>
            <n-form-item label="门店结算比例" path="orderConfirmHour" v-if="formValue.settlementType == 2">
              <n-input-group>
                <n-input-number placeholder="请输入" :min="0" :max="100" v-model:value="formValue.settlementRate" :show-button="false" style="width: 80px"/>
                <n-input-group-label>%</n-input-group-label>
              </n-input-group>
            </n-form-item>
          </n-gi>
<!--          <n-gi span="1">
            <n-form-item label="结算账户" path="settlementAccount">
              <Account ref="accountRef" :restaurantId="params.id"/>
            </n-form-item>
          </n-gi>-->
        </n-grid>
        <div style="text-align: center;">
          <n-space justify="center">
            <n-button type="info" :loading="formBtnLoading" @click="confirmForm">
              确定
            </n-button>
          </n-space>
        </div>
      </n-form>
    </n-spin>
  </div>
</template>
<script setup lang="ts">
import {onMounted, ref, watch} from "vue";
import {useProjectSettingStore} from "@/store/modules/projectSetting";
import {Edit} from "@/api/foodRestaurant";
import {useMessage} from "naive-ui";
import {useTabsViewStore} from "@/store/modules/tabsView";
import {useRouter} from "vue-router";
import {List} from "@/api/foodSettlement";
import Account from '@/views/foodRestaurant/account/index.vue';

const props = defineProps({
  formData: {
    type: Object || null,
    default: null,
  },
});
const accountRef = ref();
const settlementList = ref([]);
const emit = defineEmits(['reloadInfo']);
const router = useRouter();
const params = router.currentRoute.value.params;
const tabsViewStore = useTabsViewStore();
const show = ref(false);
const formValue = ref({
  type: 'settle',
  id: 0,
  settlementId: null,
  settlementType: 1,
  settlementRate: 0,
});
const settingStore = useProjectSettingStore();
const formBtnLoading = ref(false);
const formRef = ref<any>({});
const message = useMessage();

const rules = ref({});

function confirmForm(e) {
  e.preventDefault();
  formBtnLoading.value = true;
  formRef.value.validate((errors) => {
    if (!errors) {
      Edit(formValue.value).then((_res) => {
        message.success('操作成功');
        formBtnLoading.value = false;
        setTimeout(() => {
          if(formValue.value.id > 0){
            emit('reloadInfo');
          }else{
            // todo 关闭当前tab页，并跳转到等级列表页
            setTimeout(() => {
              tabsViewStore.closeSignal('2');
              router.push({ name: 'foodRestaurantIndex', params: {  } });
            }, 500);

          }
        }, 500);
      }).catch((err) => {
        formBtnLoading.value = false;
      });
    } else {
      formBtnLoading.value = false;
      message.error('请填写完整信息');
    }
  });
}

async function loadSettlementList(){
  let settlementArrListOrg = await List({
    status: 1,
    Pagination: false
  })
  settlementList.value = settlementArrListOrg.list
}

async function initData(data){
  formValue.value = {
    type: 'settle',
    id: data.id,
    settlementId: data.settlementId ? data.settlementId : null,
    settlementType: data.settlementType,
    settlementRate: data.settlementRate,
  }
}

watch(
  () => props.formData,
  async(value) => {
    await initData(value)
  }
);

onMounted(async() => {
  show.value = true;
  await loadSettlementList()
  await initData(props.formData)

  show.value = false;
});

</script>

<style lang="less">
</style>
