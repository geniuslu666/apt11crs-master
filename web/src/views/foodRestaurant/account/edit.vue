<template>
  <div>
    <n-modal
      v-model:show="showModal"
      :mask-closable="false"
      :show-icon="false"
      preset="dialog"
      transform-origin="center"
      :title="formValue.id > 0 ? '编辑订餐结算账户表 #' + formValue.id : '添加订餐结算账户表'"
      :style="{
        width: dialogWidth,
      }"
    >
      <n-scrollbar style="max-height: 87vh" class="pr-5">
        <n-spin :show="loading" description="请稍候...">
          <n-form
            ref="formRef"
            :model="formValue"
            :rules="rules"
            :label-placement="settingStore.isMobile ? 'top' : 'left'"
            :label-width="140"
            class="py-4"
          >
            <n-grid cols="1 s:1 m:1 l:1 xl:1 2xl:1" responsive="screen">
              <n-gi span="1">
                <n-form-item label="类型" path="type">
                  <n-radio-group v-model:value="formValue.type" name="type">
                    <n-radio-button
                      v-for="item in options.account_type"
                      :key="item.value"
                      :value="item.value"
                      :label="item.label"
                    />
                  </n-radio-group>
                </n-form-item>
              </n-gi>
              <n-gi span="1" v-if="formValue.type == 'BANKCARD'">
                <n-form-item label="开户行" path="bankName">
                  <n-input placeholder="请输入开户行" v-model:value="formValue.bankName" style="width: 250px"/>
                </n-form-item>
              </n-gi>
              <n-gi span="1" v-if="formValue.type == 'BANKCARD'">
                <n-form-item label="开户人姓名" path="bankUser">
                  <n-input placeholder="请输入开户人姓名" v-model:value="formValue.bankUser" style="width: 150px"/>
                </n-form-item>
              </n-gi>
              <n-gi span="1" v-if="formValue.type == 'BANKCARD'">
                <n-form-item label="银行账号" path="bankCard">
                  <n-input placeholder="请输入银行账号" v-model:value="formValue.bankCard" style="width: 250px"/>
                </n-form-item>
              </n-gi>
              <n-gi span="1" v-if="formValue.type == 'WECHAT'">
                <n-form-item label="微信名" path="wechatAccount">
                  <n-input placeholder="请输入微信名" v-model:value="formValue.wechatAccount" style="width: 150px"/>
                </n-form-item>
              </n-gi>
              <n-gi span="1" v-if="formValue.type == 'ALIPAY'">
                <n-form-item label="支付宝真实姓名" path="alipayName">
                  <n-input placeholder="请输入支付宝真实姓名" v-model:value="formValue.alipayName" style="width: 180px"/>
                </n-form-item>
              </n-gi>
              <n-gi span="1" v-if="formValue.type == 'ALIPAY'">
                <n-form-item label="支付宝账户" path="alipayAccount">
                  <n-input placeholder="请输入支付宝账户" v-model:value="formValue.alipayAccount" style="width: 150px"/>
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="状态" path="status">
                  <n-radio-group v-model:value="formValue.status" name="status">
                    <n-radio-button
                      v-for="item in options.sys_normal_disable"
                      :key="item.key"
                      :value="item.value"
                      :label="item.label"
                    />
                  </n-radio-group>
                </n-form-item>
              </n-gi>
            </n-grid>
          </n-form>
        </n-spin>
      </n-scrollbar>
      <template #action>
        <n-space>
          <n-button @click="closeForm">
            取消
          </n-button>
          <n-button type="info" :loading="formBtnLoading" @click="confirmForm">
            确定
          </n-button>
        </n-space>
      </template>
    </n-modal>
  </div>
</template>

<script lang="ts" setup>
import { ref, computed } from 'vue';
import { Edit, View } from '@/api/foodSettlementAccount';
import { options, State, newState } from './model';
import { useProjectSettingStore } from '@/store/modules/projectSetting';
import { useMessage, FormItemRule } from 'naive-ui';
import { adaModalWidth } from '@/utils/hotgo';

const props = defineProps({
  restaurantId: {
    type: Number,
    default: 0,
  },
});
// 表单验证规则
const rules = {
  type: {
    required: true,
    trigger: ['blur', 'change'],
    message: '请选择类型',
  },
  bankName: {
    required: true,
    trigger: ['blur', 'input'],
    validator(rule: FormItemRule, value: string) {
      if (formValue.value.type == "BANKCARD" && value.length <= 0) {
        return new Error('请输入开户行');
      }
      return true;
    },
  },
  bankUser: {
    required: true,
    trigger: ['blur', 'input'],
    validator(rule: FormItemRule, value: string) {
      if (formValue.value.type == "BANKCARD" && value.length <= 0) {
        return new Error('请输入开户人姓名');
      }
      return true;
    },
  },
  bankCard: {
    required: true,
    trigger: ['blur', 'input'],
    validator(rule: FormItemRule, value: string) {
      if (formValue.value.type == "BANKCARD" && value.length <= 0) {
        return new Error('请输入银行账号');
      }
      return true;
    },
  },
  wechatAccount: {
    required: true,
    trigger: ['blur', 'input'],
    validator(rule: FormItemRule, value: string) {
      if (formValue.value.type == "WECHAT" && value.length <= 0) {
        return new Error('请输入微信名');
      }
      return true;
    },
  },
  alipayName: {
    required: true,
    trigger: ['blur', 'input'],
    validator(rule: FormItemRule, value: string) {
      if (formValue.value.type == "ALIPAY" && value.length <= 0) {
        return new Error('请输入支付宝真实姓名');
      }
      return true;
    },
  },
  alipayAccount: {
    required: true,
    trigger: ['blur', 'input'],
    validator(rule: FormItemRule, value: string) {
      if (formValue.value.type == "ALIPAY" && value.length <= 0) {
        return new Error('请输入支付宝账户');
      }
      return true;
    },
  },
};
const emit = defineEmits(['reloadTable']);
const message = useMessage();
const settingStore = useProjectSettingStore();
const loading = ref(false);
const showModal = ref(false);
const formValue = ref<State>(newState(null));
const formRef = ref<any>({});
const formBtnLoading = ref(false);
const dialogWidth = computed(() => {
  return adaModalWidth(840);
});

function openModal(state: State) {
  showModal.value = true;

  // 新增
  if (!state || state.id < 1) {
    formValue.value = newState(state);
    formValue.value.restaurantId = props.restaurantId
    return;
  }

  // 编辑
  loading.value = true;
  View({ id: state.id })
    .then((res) => {
      formValue.value = res;
    })
    .finally(() => {
      loading.value = false;
    });
}

function confirmForm(e) {
  e.preventDefault();
  formBtnLoading.value = true;
  formRef.value.validate((errors) => {
    if (!errors) {
      Edit(formValue.value).then((_res) => {
        message.success('操作成功');
        formBtnLoading.value = false;
        setTimeout(() => {
          closeForm();
          emit('reloadTable');
        });
      }).catch((err)=>{
        formBtnLoading.value = false;
      });
    } else {
      formBtnLoading.value = false;
      message.error('请填写完整信息');
    }
  });
}

function closeForm() {
  showModal.value = false;
  loading.value = false;
}

defineExpose({
  openModal,
});
</script>

<style lang="less"></style>


