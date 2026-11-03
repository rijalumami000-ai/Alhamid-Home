<script lang="ts">
  import { onMount } from 'svelte';
  import { api } from '$lib/api';
  import type { Achievement } from '$lib/types';
  import { Trophy, Medal, Award, Calendar, Sparkles } from 'lucide-svelte';

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

<!-- Header -->
<div class="bg-[#090D16] text-white py-20 px-4 sm:px-8 relative overflow-hidden border-b border-slate-800/80">
  <div class="absolute inset-0 bg-grid-dark opacity-30 pointer-events-none"></div>
  <div class="absolute -top-24 left-1/2 -translate-x-1/2 w-[600px] h-[300px] bg-amber-500/15 blur-3xl pointer-events-none rounded-full"></div>

  <div class="max-w-6xl mx-auto space-y-4 relative z-10">
    <p class="text-xs font-mono uppercase tracking-[0.25em] font-semibold text-amber-400">
      Rekam Jejak Kejuaraan
    </p>
    <h1 class="font-display font-black text-3xl sm:text-5xl lg:text-6xl text-white tracking-tight">
      Prestasi Santri & Lembaga
    </h1>
    <p class="text-slate-300 text-base sm:text-lg max-w-2xl font-normal leading-relaxed">
      Deretan torehan emas santri Pondok Pesantren Al-Hamid dalam berbagai kompetisi musabaqah kitab kuning, sains madrasah, diplomasi bahasa, dan teknologi.
    </p>
  </div>
</div>

<div class="max-w-6xl mx-auto px-4 sm:px-8 py-20">
  <div class="grid grid-cols-1 md:grid-cols-2 gap-8">
    {#each list as item}
      <div class="bg-white rounded-3xl border border-slate-200/80 p-6 sm:p-8 shadow-sm hover:shadow-2xl hover:-translate-y-1 transition-all duration-300 flex flex-col justify-between space-y-5">
        <div class="space-y-4">
          <div class="flex items-center justify-between">
            <span class="px-3.5 py-1 rounded-full bg-amber-50 text-amber-900 border border-amber-200 text-xs font-black uppercase tracking-wider">
              {item.level}
            </span>
            <span class="text-xs text-slate-400 font-bold flex items-center gap-1.5">
              <Calendar class="w-3.5 h-3.5 text-amber-500" /> {item.year}
            </span>
          </div>

          <h2 class="font-display font-black text-xl text-slate-950 leading-snug">
            {item.title}
          </h2>

          <p class="text-slate-600 text-xs sm:text-sm leading-relaxed">
            {item.description}
          </p>
        </div>

        <div class="pt-4 border-t border-slate-100 flex items-center gap-3.5">
          <div class="w-10 h-10 rounded-full bg-teal-50 text-teal-700 flex items-center justify-center font-bold text-xs shrink-0">
            <Medal class="w-5 h-5 text-teal-600" />
          </div>
          <div>
            <div class="font-display font-bold text-xs text-slate-900">{item.recipient_name}</div>
            <div class="text-[11px] text-slate-500">Kategori: {item.category}</div>
          </div>
        </div>
      </div>
    {/each}
  </div>
</div>
