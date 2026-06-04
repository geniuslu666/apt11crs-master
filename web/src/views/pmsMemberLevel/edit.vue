<template>
  <div class="member-admin-page">
    <n-drawer v-model:show="showModal" :width="dialogWidth">
      <n-drawer-content closable :header-style="{
                    padding: '20px',
                  }" :body-content-style="{
                    padding: '20px',
                  }" :footer-style="{
                    padding: '12px 20px',
                  }">
        <template #header>
          <div style="font-weight: 500;font-size: 18px;color: #3D3D3D;line-height: 25px;">{{ formValue.id > 0 ? '编辑会员等级 #' + formValue.id : '添加会员等级' }}</div>
        </template>
        <template #footer>
          <n-button @click="closeForm" style="width: 70px;height: 35px;margin-right: 10px">
            取消
          </n-button>
          <n-button type="info" :loading="formBtnLoading" @click="confirmForm" style="width: 70px;height: 35px;">
            保存
          </n-button>
        </template>
        <n-spin :show="loading" description="请稍候...">
          <n-form :model="formValue" :rules="rules" ref="formRef"  label-placement="top"
                  label-width="auto"
                  require-mark-placement="right-hanging"
          >
            <div class="level-detail-div">
              <div class="level-detail-div-title">
                <div></div>
                基础信息
              </div>
              <n-grid :cols="1">
                <n-gi>
                  <n-form-item label="等级名称" path="levelName" >
                    <n-input placeholder="请输入会员等级名称" v-model:value="formValue.levelName" :style="{ width: '300px' }" />
                  </n-form-item>
                </n-gi>
                <n-gi>
                  <n-form-item label="等级徽章" path="levelBadge">
                    <UploadImage :maxNumber="1" v-model:value="formValue.levelBadge" />
                  </n-form-item>
                </n-gi>
                <n-gi>
                  <n-form-item label="等级卡片" path="levelCard">
                    <UploadImage :maxNumber="1" v-model:value="formValue.levelCard" />
                  </n-form-item>
                </n-gi>
                <n-gi>
                  <n-form-item label="等级字体颜色" path="wordColor">
                    <n-color-picker
                      :modes="['hex']"
                      style="width: 100px"
                      v-model:value="formValue.wordColor"
                      :show-alpha="false"
                      :swatches="['#FFFFFF', '#18A058', '#2080F0', '#F0A020', '#D03050','#000000']"
                    >
                      <template #label>
                        <div style="color: white">{{ formValue.wordColor }}</div>
                      </template>
                    </n-color-picker>
                  </n-form-item>
                </n-gi>
                <n-gi>
                  <n-form-item label="需要达到经验值" path="exp">
                    <n-input-number placeholder="请输入需要达到经验值" v-model:value="formValue.exp" :show-button="false" style="width: 300px"/>
                  </n-form-item>
                </n-gi>
                <n-gi>
                  <n-form-item label="会员等级说明" path="desc">
                    <n-input type="textarea" placeholder="请输入等级说明" v-model:value="formValue.desc" :style="{ width: '100%' }" />
                  </n-form-item>
                </n-gi>
              </n-grid>
            </div>
            <div class="level-detail-div">
              <div class="level-detail-div-title">
                <div></div>
                权益
              </div>
              <n-grid :cols="2">
                <n-gi>
                  <n-form-item label="酒店场景获取积分倍率" path="pointFeedback">
                    <n-input-number placeholder="请输入酒店场景获取积分倍率" v-model:value="formValue.hotelGetRate" :show-button="false" style="width: 290px">
                      <template #suffix>
                        倍
                      </template>
                    </n-input-number>
                  </n-form-item>
                </n-gi>
                <n-gi>
                  <n-form-item label="餐饮场景获取积分倍率" path="consumeDiscount">
                    <n-input-number placeholder="请输入餐饮场景获取积分倍率" v-model:value="formValue.foodGetRate" :show-button="false" style="width: 290px">
                      <template #suffix>
                        倍
                      </template>
                    </n-input-number>
                  </n-form-item>
                </n-gi>
                <n-gi>
                  <n-form-item label="按摩场景获取积分倍率" path="spaGetRate">
                    <n-input-number placeholder="请输入按摩场景获取积分倍率" v-model:value="formValue.spaGetRate" :show-button="false" style="width: 290px">
                      <template #suffix>
                        倍
                      </template>
                    </n-input-number>
                  </n-form-item>
                </n-gi>
                <n-gi>
                  <n-form-item label="接送机/包车场景获取积分倍率" path="carGetRate">
                    <n-input-number placeholder="请输入接送机/包车场景获取积分倍率" v-model:value="formValue.carGetRate" :show-button="false" style="width: 290px">
                      <template #suffix>
                        倍
                      </template>
                    </n-input-number>
                  </n-form-item>
                </n-gi>
                <n-gi>
                  <n-form-item label="储物柜场景获取积分倍率" path="cabinetGetRate">
                    <n-input-number placeholder="请输入储物柜场景获取积分倍率" v-model:value="formValue.cabinetGetRate" :show-button="false" style="width: 290px">
                      <template #suffix>
                        倍
                      </template>
                    </n-input-number>
                  </n-form-item>
                </n-gi>
              </n-grid>
            </div>
          </n-form>
        </n-spin>
      </n-drawer-content>
    </n-drawer>

  </div>
</template>

<script lang="ts" setup>
import {computed, ref} from "vue";
import {adaModalWidth} from "@/utils/hotgo";
import {newState, rules, State} from "@/views/pmsMemberLevel/model";
import {Edit, View} from "@/api/pmsMemberLevel";
import {useMessage} from "naive-ui";

const emit = defineEmits(['reloadTable']);
const loading = ref(false);
const showModal = ref(false);
const formRef: any = ref(null);
const formValue = ref<State>(newState(null));
const formBtnLoading = ref(false);
const message = useMessage();

const dialogWidth = computed(() => {
  return adaModalWidth(650);
});

function confirmForm(e) {
  e.preventDefault();
  formBtnLoading.value = true;
  formRef.value.validate((errors) => {
    if (!errors) {
      Edit(formValue.value).then((_res) => {
        message.success('操作成功');
        formBtnLoading.value = false;
        closeForm();
        emit('reloadTable');
      }).catch((err) => {
        formBtnLoading.value = false;
      });
    } else {
      formBtnLoading.value = false;
      message.error('请填写完整信息');
    }
  });
}

function openModal(state: State) {
  showModal.value = true;
  if(state.id > 0){
    formValue.value.id = state.id
    loading.value = true;
    View({ id: state.id })
      .then((res) => {
        formValue.value = res;
      })
      .finally(() => {
        loading.value = false;
      });
  }
}

function closeForm() {
  showModal.value = false;
  loading.value = false;
}

defineExpose({
  openModal,
});
</script>
<style lang="less" scoped>
.level-detail-div{
  &-title{
    display: flex;
    align-items: center;
    font-weight: 500;
    font-size: 16px;
    color: #3D3D3D;
    line-height: 22px;
    margin-bottom: 15px;
    div{
      margin-right: 5px;
      width: 6px;
      height: 15px;
      background: #053DC8;
    }
  }
  &-item{
    &-title{
      font-weight: 400;
      font-size: 14px;
      color: #3D3D3D;
      line-height: 20px;
      margin-bottom: 8px;
    }
    &-content{

    }
  }
}
</style>
