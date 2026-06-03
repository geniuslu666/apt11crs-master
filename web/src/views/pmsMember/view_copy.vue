<template>
  <div>
    <n-drawer v-model:show="showModal" :width="dialogWidth">
      <n-drawer-content title="会员信息详情" closable>
        <n-spin :show="loading" description="请稍候...">
          <div style="width: 100%">
            <n-grid
              x-gap="10"
              :cols="3"
              class="f12 mb-5"
              style="border-bottom: 1px solid #eeeeeeff; line-height: 25px"
            >
              <n-gi>
                <div class="cblue f14 fw flex-row">
                  <img
                    v-if="formValue.avatar"
                    :src="formValue.avatar"
                    style="width: 40px; height: 40px; border-radius: 50%; object-fit: cover"
                  />
                  <img
                    v-else
                    src="~@/assets/images/mrtx.png"
                    style="width: 40px; height: 40px; border-radius: 50%"
                  />
                  <div class="flex-item ml-2" style="line-height: 20px">
                    {{ formValue.memberNo }}
                    <a class="c999 f12 ml-1" v-copy="formValue.memberNo">复制</a>
                    <div>
                      <span class="cwarning">level：{{ formValue.level }}</span>
                    </div>
                  </div>
                </div>
              </n-gi>
              <n-gi>
                <div class="c999">注册来源</div>
                <div class="c333"> {{ formValue.source }}</div>
              </n-gi>
              <n-gi>
                <div class="c999">积分/经验</div>
                <div class="c333 mb-3">
                  <span class="csuccess">{{ formValue.balance }}</span> /
                  <span class="c999">{{ formValue.exp }}</span>
                </div>
              </n-gi>

              <n-gi>
                <div class="c999">名</div>
                <div class="c333">{{ formValue.firstName }}</div>
              </n-gi>
              <n-gi>
                <div class="c999">姓</div>
                <div class="c333">{{ formValue.lastName }}</div>
              </n-gi>
              <n-gi>
                <div class="c999">全名</div>
                <div class="c333 mb-3">{{ formValue.fullName }}</div>
              </n-gi>
            </n-grid>
            <n-grid x-gap="10" :cols="2" class="f12" style="line-height: 30px">
              <n-gi>
                <div class="c999">手机号</div>
                <div class="c333"
                  >{{ formValue.phoneArea }} {{ formValue.phone }}
                  <a class="c999 f12 ml-1" v-if="formValue.phone" copy="formValue.phone">复制</a></div
                >
              </n-gi>
              <n-gi>
                <div class="c999">邮箱</div>
                <div class="c333 mb-3"
                  >{{ formValue.mail }}
                  <span class="c999 f12 ml-1"  v-if="formValue.mail"  v-copy="formValue.mail">复制</span></div
                >
              </n-gi>

              <n-gi>
                <div class="c999">生日</div>
                <div class="c333">{{ formValue.birthday }}</div>
              </n-gi>
              <n-gi>
                <div class="c999">推荐人</div>
                <div class="c333 mb-3">{{ formValue.referrer }}</div>
              </n-gi>

              <n-gi>
                <div class="c999">上次登录ip</div>
                <div class="c333 mb-3">{{ formValue.lastLoginIp }}</div>
              </n-gi>

              <n-gi>
                <div class="c999">上次登录时间</div>
                <div class="c333">{{ formValue.lastLogin }}</div>
              </n-gi>

              <n-gi>
                <div class="c999">注册ip</div>
                <div class="c333 mb-3">{{ formValue.registerIp }}</div>
              </n-gi>

              <n-gi>
                <div class="c999">注册时间</div>
                <div class="c333">{{ formValue.registerTime }}</div>
              </n-gi>

              <n-gi>
                <div class="c999">国籍</div>
                <div class="c333 mb-3">{{ formValue.nationality }}</div>
              </n-gi>
            </n-grid>
            <n-grid :cols="1">
              <n-gi>
                <div class="c999">地址</div>
                <div class="c333">{{ formValue.address }}</div>
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
  import { useMessage } from 'naive-ui';
  import { View } from '@/api/pmsMember';
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
