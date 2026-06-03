<template>
  <div>
    <n-drawer v-model:show="showModal" :width="dialogWidth" :z-index="99">
      <n-drawer-content closable :header-style="{ padding: '20px' }" :body-content-style="{ padding: '20px' }" :footer-style="{ padding: '12px 20px' }">
        <template #header>
          <div style="font-weight: 500;font-size: 18px;color: #3D3D3D;line-height: 25px;">
            {{ formValue.id > 0 ?  '编辑车型 #' + formValue.id : '新增车型' }}
<!--            {{ formValue.id > 0 ? (productName ? productName + '-编辑车型 #' + formValue.id : '编辑车型 #' + formValue.id) : (productName ? productName + '-新增车型' : '添加车型') }}-->
          </div>
        </template>
        <template #footer>
          <n-button @click="closeForm" style="width: 70px;height: 35px;margin-right: 10px">取消</n-button>
          <n-button type="info" :loading="formBtnLoading" @click="confirmForm" style="width: 70px;height: 35px;">保存</n-button>
        </template>
        <n-spin :show="loading" description="请稍候...">
          <n-form ref="formRef" :model="formValue" :rules="rules" label-placement="top" label-width="auto">
            <n-grid :cols="2">
              <n-gi>
                <n-form-item label="车型名称_简体中文" path="nameZh">
                  <n-input placeholder="简体中文" v-model:value="nameLanguage.zh" style="width: 300px" />
                </n-form-item>
              </n-gi>
              <n-gi>
                <n-form-item label="车型名称_English" path="nameEn">
                  <n-input placeholder="English" v-model:value="nameLanguage.en" style="width: 300px" />
                </n-form-item>
              </n-gi>
              <n-gi>
                <n-form-item label="车型名称_日本語" path="nameJa">
                  <n-input placeholder="日本語" v-model:value="nameLanguage.ja" style="width: 300px" />
                </n-form-item>
              </n-gi>
              <n-gi>
                <n-form-item label="车型名称_한국어" path="nameKo">
                  <n-input placeholder="한국어" v-model:value="nameLanguage.ko" style="width: 300px" />
                </n-form-item>
              </n-gi>
              <n-gi span="2">
                <n-form-item label="车型名称_繁体中文" path="nameZhCN">
                  <n-input placeholder="繁体中文" v-model:value="nameLanguage.zh_CN" style="width: 300px" />
                </n-form-item>
              </n-gi>
              <n-gi>
                <n-form-item label="价格" path="price">
                  <n-input-group>
                    <n-input-number placeholder="请输入价格" :min="0" v-model:value="formValue.price" style="width: 150px" />
                    <n-input-group-label>JPY</n-input-group-label>
                  </n-input-group>
                </n-form-item>
              </n-gi>
              <n-gi>
                <n-form-item label="每日接待人数" path="dailyCapacity">
                  <n-input-group>
                    <n-input-number placeholder="请输入人数" :min="1" v-model:value="formValue.dailyCapacity" style="width: 150px" />
                    <n-input-group-label>人</n-input-group-label>
                  </n-input-group>
                </n-form-item>
              </n-gi>
              <n-gi span="2">
                <n-form-item label="集合地点" path="meetingPlace">
                  <n-input v-model:value="formValue.meetingPlace" placeholder="请输入集合地点" style="width: 600px"/>
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="集合地点纬度" path="ggLat">
                  <n-input
                    placeholder="请输入纬度"
                    v-model:value="formValue.ggLat"
                    style="width: 250px"
                  />
                </n-form-item>
              </n-gi>
              <n-gi span="1 600:2">
                <n-form-item label="集合地点经度" path="ggLng">
                  <n-input
                    placeholder="请输入经度"
                    v-model:value="formValue.ggLng"
                    style="width: 250px"
                  />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="集合时间" path="meetingTime">
                  <n-time-picker placeholder="请选择" format="HH:mm" v-model:formatted-value="formValue.meetingTime" />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="联系电话" path="contactMobile">
                  <n-input v-model:value="formValue.contactMobile" placeholder="请输入联系电话" style="width: 250px" />
                </n-form-item>
              </n-gi>
              <n-gi>
                <n-form-item label="状态" path="status">
                  <n-radio-group v-model:value="formValue.status" name="status">
                    <n-radio-button :value="1" label="启用" />
                    <n-radio-button :value="2" label="禁用" />
                  </n-radio-group>
                </n-form-item>
              </n-gi>
              <n-gi>
                <n-form-item label="排序">
                  <n-input-number v-model:value="formValue.sort" :min="0" placeholder="越大越靠前" style="width: 100px" />
                </n-form-item>
              </n-gi>
              <n-gi span="2">
                <n-form-item label="是否允许取消" path="allowCancel">
                  <n-radio-group v-model:value="formValue.allowCancel" name="allowCancel">
                    <n-space>
                      <n-radio :value="1">
                        允许
                      </n-radio>
                      <n-radio :value="2">
                        不允许
                      </n-radio>
                    </n-space>
                  </n-radio-group>
                </n-form-item>
                <n-form-item label="取消政策" path="afterConfirmCancel" v-if="formValue.allowCancel == 1" >
                  <n-input-group>
                    <n-input-group-label>距离订单服务开始</n-input-group-label>
                    <n-input-number placeholder="请输入" :min="0" v-model:value="formValue.freeCancelHours" style="width: 100px" />
                    <n-input-group-label>小时前可免费取消，之后取消费为订单金额的</n-input-group-label>
                    <n-input-number placeholder="请输入" :min="0" :max="100" v-model:value="formValue.cancelFeePercent" style="width: 100px" />
                    <n-input-group-label>%</n-input-group-label>
                  </n-input-group>
                </n-form-item>
              </n-gi>
            </n-grid>
          </n-form>
        </n-spin>
      </n-drawer-content>
    </n-drawer>
  </div>
</template>

<script lang="ts" setup>
import { ref, computed, reactive } from 'vue';
import { useMessage } from 'naive-ui';
import { Edit, View } from '@/api/travelProductSku';
import { State, newState, rules } from './model';
import { adaModalWidth } from '@/utils/hotgo';
import { jsontoobj } from '@/utils/smjcomm';
// import {View as ProductView} from "@/api/travelProduct";

const emit = defineEmits(['reloadTable']);
const message = useMessage();
const loading = ref(false);
const showModal = ref(false);
const formBtnLoading = ref(false);
const formRef = ref<any>({});
const formValue = ref<State>(newState(null));
const initProductId = ref<number>(0);
const dialogWidth = computed(() => adaModalWidth(700));

const nameLanguage = reactive({ zh: '', en: '', ja: '', ko: '', zh_CN: '' });
// const productName = ref('');

function confirmForm(e) {
  const hasName = nameLanguage.zh || nameLanguage.en || nameLanguage.ja || nameLanguage.ko || nameLanguage.zh_CN;
  if (!hasName) {
    message.warning('车型名称至少填写一种语言');
    return;
  }
  e.preventDefault();
  formBtnLoading.value = true;
  formRef.value?.validate((errors) => {
    if (!errors) {
      formValue.value.nameLanguage = nameLanguage;
      Edit(formValue.value).then(() => {
        message.success('操作成功');
        formBtnLoading.value = false;
        closeForm();
        emit('reloadTable');
      }).catch(() => {
        formBtnLoading.value = false;
      });
    } else {
      message.error('请填写完整信息');
      formBtnLoading.value = false;
    }
  });
}

function closeForm() {
  showModal.value = false;
  loading.value = false;
}

// async function loadProductView(){
//   let info = await ProductView({ id: initProductId.value })
//   productName.value = info.title
// }

async function openModal(state: State, productId: number) {
  showModal.value = true;
  loading.value = true;
  initProductId.value = productId;

  // await loadProductView()

  if (!state || state.id < 1) {
    formValue.value = newState(state);
    formValue.value.productId = productId;
    nameLanguage.zh = '';
    nameLanguage.en = '';
    nameLanguage.ja = '';
    nameLanguage.ko = '';
    nameLanguage.zh_CN = '';
    loading.value = false;
    return;
  }

  const res = await View({ id: state.id, isLanguage: true });
  formValue.value = res;

  if(formValue.value.meetingTime == ''){
    formValue.value.meetingTime = null
  }

  // console.log(formValue.value)

  if (res.nameLanguage) {
    const nl = jsontoobj(res.nameLanguage);
    nameLanguage.zh = nl.zh?.content || '';
    nameLanguage.en = nl.en?.content || '';
    nameLanguage.ja = nl.ja?.content || '';
    nameLanguage.ko = nl.ko?.content || '';
    nameLanguage.zh_CN = nl.zh_CN?.content || nl.zh_cn?.content || '';
  }
  loading.value = false;
}

defineExpose({ openModal });
</script>
