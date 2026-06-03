<template>
  <div>
    <n-drawer v-model:show="showModal" :width="dialogWidth">
      <n-drawer-content closable :header-style="{
                    padding: '20px',
                  }" :body-content-style="{
                    padding: '25px 0 0',
                  }">
        <template #header>
          <div style="font-weight: 500;font-size: 18px;color: #3D3D3D;line-height: 25px;">Banner详情</div>
        </template>
        <n-spin :show="loading" description="请稍候...">
          <div class="level-detail-div" style="padding: 0 20px;">
            <n-grid y-gap="25" :cols="2">
              <n-gi span="2">
                <div class="level-detail-div-item">
                  <div class="level-detail-div-item-title">名称</div>
                  <div class="level-detail-div-item-content">
                    {{ formValue.bannerName }}
                  </div>
                </div>
              </n-gi>
              <n-gi span="2">
                <div class="level-detail-div-item">
                  <div class="level-detail-div-item-title">轮播图</div>
                  <div class="level-detail-div-item-content">
                    <n-image style="width: 300px;border-radius: 2px" :src="formValue.bannerImage"/>
                  </div>
                </div>
              </n-gi>
              <n-gi>
                <div class="level-detail-div-item">
                  <div class="level-detail-div-item-title">语言</div>
                  <div class="level-detail-div-item-content1" :class="getOptionTag(options.language, formValue?.language)">{{ getOptionLabel(options.language, formValue?.language) }}</div>
                </div>
              </n-gi>
              <n-gi>
                <div class="level-detail-div-item">
                  <div class="level-detail-div-item-title">模块</div>
                  <div class="level-detail-div-item-content1" :class="getOptionTag(options.model, formValue?.model)">{{ getOptionLabel(options.model, formValue?.model) }}</div>
                </div>
              </n-gi>
              <n-gi>
                <div class="level-detail-div-item">
                  <div class="level-detail-div-item-title">内外联</div>
                  <div class="level-detail-div-item-content1" :class="getOptionTag(options.chain, formValue?.chain)">{{ getOptionLabel(options.chain, formValue?.chain) }}</div>
                </div>
              </n-gi>
              <n-gi>
                <div class="level-detail-div-item">
                  <div class="level-detail-div-item-title">状态</div>
                  <div class="level-detail-div-item-content1" :class="getOptionTag(options.sys_normal_disable, formValue?.bannerStatus)">{{ getOptionLabel(options.sys_normal_disable, formValue?.bannerStatus) }}</div>
                </div>
              </n-gi>
              <n-gi span="2">
                <div class="level-detail-div-item">
                  <div class="level-detail-div-item-title">链接内容</div>
                  <div class="level-detail-div-item-content">
                    {{ formValue.path }}
                  </div>
                </div>
              </n-gi>
            </n-grid>
          </div>
        </n-spin>
      </n-drawer-content>
    </n-drawer>
  </div>
</template>

<script lang="ts" setup>
import { computed, ref } from 'vue';
import { View } from '@/api/pmsBanner';
import { State, newState, options } from './model';
import { adaModalWidth, getOptionLabel, getOptionTag } from '@/utils/hotgo';

const loading = ref(false);
const showModal = ref(false);
const formValue = ref(newState(null));
const dialogWidth = computed(() => {
  return adaModalWidth(650);
});


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

<style lang="less" scoped>
.level-detail-div{
  margin-bottom: 25px;
  &-title{
    display: flex;
    align-items: center;
    font-weight: 500;
    font-size: 16px;
    color: #3D3D3D;
    line-height: 22px;
    margin-bottom: 25px;
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
      font-size: 12px;
      color: #707070;
      line-height: 17px;
      margin-bottom: 8px;
    }
    &-content{
      font-weight: 500;
      font-size: 14px;
      color: #3D3D3D;
      line-height: 20px;
      display: flex;
      align-items: center;
      .color{
        width: 20px;
        height: 20px;
        margin-right: 8px;
        border-radius: 2px;
      }
    }
    &-content1{
      padding: 0 6px;
      height: 24px;
      line-height: 24px;
      font-weight: 400;
      font-size: 14px;
      display: inline-block;
      &.success{
        color: #19A158;
        background: #E3F4EB;
      }
      &.info{
        color: #3F9EFF;
        background: #ECF5FF;
      }
      &.default{
        color: #919399;
        background: #F4F4F5;
      }
      &.warning{
        color: #EFA020;
        background: #FDF1DD;
      }
      &.error{
        color: #F56C6C;
        background: #FEF0F0;
      }
    }
  }
}
</style>


