<template>
  <!-- Hero Section -->
  <header class="relative pt-32 pb-20 px-4 overflow-hidden"
          :style="{ backgroundColor: 'var(--header-bg)' }">
    <div class="absolute inset-0 pointer-events-none">
      <div class="absolute -top-[10%] left-1/4 w-[40%] h-[40%] rounded-full blur-[120px]"
            :style="{ backgroundColor: 'var(--glow-orange)' }"></div>
      <div class="absolute -bottom-[10%] right-1/4 w-[40%] h-[40%] rounded-full blur-[120px]"
            :style="{ backgroundColor: 'var(--glow-amber)' }"></div>
    </div>

    <div class="relative z-10 text-center max-w-4xl mx-auto">
      <span class="inline-block py-1 px-3 rounded-full text-sm font-bold mb-6 tracking-wide uppercase border"
            :style="{
              backgroundColor: 'var(--badge-bg)',
              borderColor: 'var(--badge-border)',
              color: 'var(--accent)'
            }">
        Get In Touch
      </span>
      <h1 class="text-5xl md:text-6xl font-extrabold mb-6 leading-tight"
          :style="{ color: 'var(--header-text)' }">
        We're Here to <span :style="{ color: 'var(--accent)' }">Help</span>
      </h1>
      <p class="text-lg md:text-xl mb-10 max-w-2xl mx-auto"
          :style="{ color: 'var(--header-muted)' }">
        Have a question? We'd love to hear from you.
      </p>
    </div>
  </header>

  <!-- Main Content -->
  <main class="py-20 px-4 sm:px-6 lg:px-8"
        :style="{ backgroundColor: 'var(--surface)' }">
    <div class="max-w-4xl mx-auto">

      <div class="grid grid-cols-1 lg:grid-cols-2 gap-12">
        <!-- Contact Form -->
        <div>
          <h2 class="text-3xl font-bold mb-8"
              :style="{ color: 'var(--text-color)' }">
            Send Us a Message
          </h2>

          <div class="p-8 rounded-2xl border"
                :style="{
                  backgroundColor: 'var(--card-bg)',
                  borderColor: 'var(--border)'
                }">
            <form class="space-y-6">
              <div>
                <label class="block text-sm font-medium mb-2"
                        :style="{ color: 'var(--text-color)' }">
                  Name
                </label>
                <input type="text"
                        class="w-full px-4 py-3 rounded-lg border focus:outline-none focus:ring-2 focus:ring-violet-500"
                        :style="{
                          backgroundColor: 'var(--bg-color)',
                          borderColor: 'var(--border)',
                          color: 'var(--text-color)'
                        }">
              </div>

              <div>
                <label class="block text-sm font-medium mb-2"
                        :style="{ color: 'var(--text-color)' }">
                  Email
                </label>
                <input type="email"
                        class="w-full px-4 py-3 rounded-lg border focus:outline-none focus:ring-2 focus:ring-violet-500"
                        :style="{
                          backgroundColor: 'var(--bg-color)',
                          borderColor: 'var(--border)',
                          color: 'var(--text-color)'
                        }">
              </div>

              <div>
                <label class="block text-sm font-medium mb-2"
                        :style="{ color: 'var(--text-color)' }">
                  Message
                </label>
                <textarea rows="4"
                          class="w-full px-4 py-3 rounded-lg border focus:outline-none focus:ring-2 focus:ring-violet-500"
                          :style="{
                            backgroundColor: 'var(--bg-color)',
                            borderColor: 'var(--border)',
                            color: 'var(--text-color)'
                          }"></textarea>
              </div>

              <button type="submit"
                      class="w-full px-6 py-3 rounded-lg font-semibold transition-all transform hover:scale-105 shadow-lg text-white"
                      :style="{ backgroundColor: 'var(--accent)' }">
                Send Message
              </button>
            </form>
          </div>
        </div>

        <!-- Contact Info -->
        <div>
          <h2 class="text-3xl font-bold mb-8"
              :style="{ color: 'var(--text-color)' }">
            Other Ways to Connect
          </h2>

          <div class="space-y-6">
            <div class="p-6 rounded-xl border hover:-translate-y-1 transition-transform"
                  :style="{
                    backgroundColor: 'var(--card-bg)',
                    borderColor: 'var(--border)'
                  }">
              <div class="flex items-start gap-4">
                <EnvelopeIcon class="w-8 h-8 flex-shrink-0" :style="{ color: 'var(--accent)' }" />
                <div>
                  <h3 class="text-lg font-bold mb-1"
                      :style="{ color: 'var(--text-color)' }">
                    Email
                  </h3>
                  <p :style="{ color: 'var(--muted)' }">
                    hello@example.com
                  </p>
                </div>
              </div>
            </div>

            <div class="p-6 rounded-xl border hover:-translate-y-1 transition-transform"
                  :style="{
                    backgroundColor: 'var(--card-bg)',
                    borderColor: 'var(--border)'
                  }">
              <div class="flex items-start gap-4">
                <DevicePhoneMobileIcon class="w-8 h-8 flex-shrink-0" :style="{ color: 'var(--accent)' }" />
                <div>
                  <h3 class="text-lg font-bold mb-1"
                      :style="{ color: 'var(--text-color)' }">
                    Social Media
                  </h3>
                  <p :style="{ color: 'var(--muted)' }">
                    Follow us for updates
                  </p>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- Footer Section -->
      <div class="mt-20 pt-12 border-t"
            :style="{ borderColor: 'var(--border)' }">
        <div class="grid grid-cols-1 md:grid-cols-4 gap-8">
          <div>
            <span class="font-bold text-2xl tracking-wider uppercase mb-4 inline-block"
                  :style="{ color: 'var(--text-color)' }">
              App
            </span>
          </div>
          <div>
            <h3 class="font-bold mb-4"
                :style="{ color: 'var(--text-color)' }">Explore</h3>
            <ul class="space-y-2">
              <li><router-link to="/" class="hover:text-violet-500 transition"
                              :style="{ color: 'var(--muted)' }">Home</router-link></li>
              <li><router-link to="/about" class="hover:text-violet-500 transition"
                              :style="{ color: 'var(--muted)' }">About</router-link></li>
            </ul>
          </div>
          <div>
            <h3 class="font-bold mb-4"
                :style="{ color: 'var(--text-color)' }">Connect</h3>
            <ul class="space-y-2">
              <li><a href="#" class="hover:text-violet-500 transition"
                    :style="{ color: 'var(--muted)' }">Social</a></li>
            </ul>
          </div>
        </div>
      </div>
    </div>
  </main>
</template>

<script setup lang="ts">
import { EnvelopeIcon, DevicePhoneMobileIcon } from '@heroicons/vue/24/outline'
</script>
