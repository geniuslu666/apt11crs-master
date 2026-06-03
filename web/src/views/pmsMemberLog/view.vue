<template>
  <div>
    <n-drawer v-model:show="showModal" :width="dialogWidth">
      <n-drawer-content closable :header-style="{
                    padding: '20px',
                  }" :body-content-style="{
                    padding: '25px 20px',
                  }">
        <template #header>
          <div style="font-weight: 500;font-size: 18px;color: #3D3D3D;line-height: 25px;">会员登录日志详情</div>
        </template>
        <n-spin :show="loading" description="请稍候...">
          <n-grid y-gap="25" :cols="2">
            <n-gi>
              <div class="content-detail-item">
                <div class="content-detail-item-title">会员ID</div>
                <div class="content-detail-item-content">{{ formValue.memberId }}</div>
              </div>
            </n-gi>
            <n-gi>
              <div class="content-detail-item">
                <div class="content-detail-item-title">登录方式</div>
                <div>
                  <div class="content-detail-item-content1" :class="getOptionTag(options.login_type, formValue?.loginType)">{{ getOptionLabel(options.login_type, formValue?.loginType) }}</div>
                </div>
              </div>
            </n-gi>
            <n-gi>
              <div class="content-detail-item">
                <div class="content-detail-item-title">登录时间</div>
                <div class="content-detail-item-content">{{ formValue.loginTime }}</div>
              </div>
            </n-gi>
            <n-gi>
              <div class="content-detail-item">
                <div class="content-detail-item-title">登录IP</div>
                <div class="content-detail-item-content">{{ formValue.loginIp }}</div>
              </div>
            </n-gi>
            <n-gi>
              <div class="content-detail-item">
                <div class="content-detail-item-title">过期时间</div>
                <div class="content-detail-item-content">{{ formValue.expirTime }}</div>
              </div>
            </n-gi>
            <n-gi>
              <div class="content-detail-item">
                <div class="content-detail-item-title">登录token</div>
                <div class="content-detail-item-token">
                  <div>{{formValue.token}}</div>
                  <a href="javascript:void(0);" @click="copyToken">复制</a>
                </div>
              </div>
            </n-gi>
          </n-grid>
        </n-spin>
      </n-drawer-content>
    </n-drawer>
  </div>
</template>

<script lang="ts" setup>
  import { computed, ref } from 'vue';
  import { useMessage } from 'naive-ui';
  import { View } from '@/api/pmsMemberLog';
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

  function copyToken(){
    navigator.clipboard.writeText(formValue.value.token).then(() => {
      message.success('复制成功')
    }).catch(() => {
      message.error('复制失败')
    })
  }

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

<style lang="less" scoped>
.content-detail-item{
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
  }
  &-token{
    display: flex;
    align-items: center;
    div{
      width: 230px;
      height: 30px;
      line-height: 30px;
      padding: 0 8px;
      white-space: nowrap; /* 防止文本换行 */
      overflow: hidden; /* 隐藏溢出的内容 */
      text-overflow: ellipsis; /* 使用省略号表示溢出的文本 */
      background: #FAFAFC;
      border: 1px solid #E0E0E6;
      border-radius: 2px;
      font-weight: 500;
      font-size: 14px;
      color: #C2C2C2;
    }
    a{
      display: block;
      margin-left: 10px;
      width: 60px;
      height: 30px;
      line-height: 30px;
      background: #053DC8;
      border-radius: 2px;
      text-align: center;
      font-weight: 400;
      font-size: 14px;
      color: #FFFFFF;
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

</style>
