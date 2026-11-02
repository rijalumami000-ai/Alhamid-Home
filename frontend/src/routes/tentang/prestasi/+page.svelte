<script lang="ts">
  import { onMount } from 'svelte';
  import { api } from '$lib/api';
  import type { Achievement } from '$lib/types';
  import { Trophy, Medal, Award, Calendar } from 'lucide-svelte';

  let list = $state<Achievement[]>([]);

  onMount(async () => {
    try {
      list = await api.getAchievements();
    } catch (e) {
      console.error(e);
    }
  });
</script>

<svelte:head>
  <title>Prestasi Santri & Lembaga | Pondok Pesantren Al-Hamid</title>
</svelte:head>

<div class="bg-gradient-to-r from-emerald-950 via-emerald-900 to-emerald-800 text-white py-16 px-4 sm:px-8">
  <div class="max-w-6xl mx-auto space-y-4">
    <div class="inline-flex items-center gap-2 px-3.5 py-1.5 rounded-full bg-amber-500/20 text-amber-300 text-xs font-semibold border border-amber-500/30">
      Capaian Gemilang
    </div>
    <h1 class="font-serif-luxury font-bold text-3xl sm:text-5xl text-white">
      Prestasi Santri & Lembaga
    </h1>
    <p class="text-emerald-100/80 text-base max-w-2xl">
      Deretan torehan emas santri Pondok Pesantren Al-Hamid dalam berbagai kompetisi keagamaan, sains, bahasa, dan seni.
    </p>
  </div>
</div>

<div class="max-w-6xl mx-auto px-4 sm:px-8 py-16">
  <div class="grid grid-cols-1 md:grid-cols-2 gap-8">
    {#each list as item}
      <div class="bg-white rounded-3xl border border-slate-100 p-6 sm:p-8 shadow-sm hover:shadow-xl transition-all duration-300 flex flex-col justify-between space-y-4">
        <div class="space-y-3">
          <div class="flex items-center justify-between">
            <span class="px-3 py-1 rounded-full bg-amber-50 text-amber-800 border border-amber-200 text-xs font-bold uppercase tracking-wide">
              {item.level}
            </span>
            <span class="text-xs text-slate-400 font-semibold flex items-center gap-1">
              <Calendar class="w-3.5 h-3.5" /> {item.year}
            </span>
          </div>

          <h2 class="font-serif-luxury font-bold text-xl text-slate-900 leading-snug">
            {item.title}
          </h2>

          <p class="text-slate-600 text-xs sm:text-sm leading-relaxed">
            {item.description}
          </p>
        </div>

        <div class="pt-4 border-t border-slate-100 flex items-center gap-3">
          <div class="w-9 h-9 rounded-full bg-emerald-100 text-emerald-800 flex items-center justify-center font-bold text-xs shrink-0">
            <Medal class="w-4 h-4" />
          </div>
          <div>
            <div class="text-xs font-semibold text-slate-800">{item.recipient_name}</div>
            <div class="text-[11px] text-slate-400">Kategori: {item.category}</div>
          </div>
        </div>
      </div>
    {/each}
  </div>
</div>

