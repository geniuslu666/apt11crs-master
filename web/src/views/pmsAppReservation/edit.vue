<template>
  <div>
    <n-modal
      v-model:show="showModal"
      :mask-closable="false"
      :show-icon="false"
      preset="dialog"
      transform-origin="center"
      :title="formValue.id > 0 ? '编辑系统入住订单 #' + formValue.id : '添加系统入住订单'"
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
            :label-width="100"
            class="py-4"
          >
            <n-grid cols="1 s:1 m:1 l:1 xl:1 2xl:1" responsive="screen">
              <n-gi span="1">
                <n-form-item label="房间预订ID" path="uid">
                  <n-input placeholder="请输入房间预订ID" v-model:value="formValue.uid" />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="用户ID" path="memberId">
                  <n-input-number placeholder="请输入用户ID" v-model:value="formValue.memberId" />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="物业ID" path="puid">
                  <n-input placeholder="请输入物业ID" v-model:value="formValue.puid" />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="系统订单号" path="orderSn">
                  <n-input placeholder="请输入系统订单号" v-model:value="formValue.orderSn" />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="票号" path="ticketId">
                  <n-input placeholder="请输入票号" v-model:value="formValue.ticketId" />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="房型信息" path="roomType">
                  <n-input placeholder="请输入房型信息" v-model:value="formValue.roomType" />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="房间单元" path="roomUnit">
                  <n-input placeholder="请输入房间单元" v-model:value="formValue.roomUnit" />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="预订人信息" path="guestProfile">
                  <n-input placeholder="请输入预订人信息" v-model:value="formValue.guestProfile" />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="入住日期" path="checkinDate">
                  <DatePicker v-model:formValue="formValue.checkinDate" type="date" />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="退房日期" path="checkoutDate">
                  <DatePicker v-model:formValue="formValue.checkoutDate" type="date" />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="入住时间" path="checkinTime">
                  <n-input placeholder="请输入入住时间" v-model:value="formValue.checkinTime" />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="退房时间" path="checkoutTime">
                  <n-input placeholder="请输入退房时间" v-model:value="formValue.checkoutTime" />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="预订状态" path="status">
                  <n-select v-model:value="formValue.status" :options="options.status" />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="入住状态" path="checkinStatus">
                  <n-select v-model:value="formValue.checkinStatus" :options="options.checkin_status" />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="支付状态" path="orderStatus">
                  <n-select v-model:value="formValue.orderStatus" :options="options.order_status" />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="成人数量" path="adultCount">
                  <n-input-number placeholder="请输入成人数量" v-model:value="formValue.adultCount" />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="儿童数量" path="childCount">
                  <n-input-number placeholder="请输入儿童数量" v-model:value="formValue.childCount" />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="预订费" path="bookingFee">
                  <n-input-number placeholder="请输入预订费" v-model:value="formValue.bookingFee" />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="渠道费" path="channelFee">
                  <n-input-number placeholder="请输入渠道费" v-model:value="formValue.channelFee" />
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
  import { Edit, View } from '@/api/pmsAppReservation';
  import { options, State, newState, rules } from './model';
  import DatePicker from '@/components/DatePicker/datePicker.vue';
  import { useProjectSettingStore } from '@/store/modules/projectSetting';
  import { useMessage } from 'naive-ui';
  import { adaModalWidth } from '@/utils/hotgo';

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
          setTimeout(() => {
            closeForm();
            emit('reloadTable');
          });
        });
      } else {
        message.error('请填写完整信息');
      }
      formBtnLoading.value = false;
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