<script lang="ts">
  import { onMount } from 'svelte';
  import { api } from '$lib/api';
  import type { Alumni } from '$lib/types';
  import { Users, MapPin, Briefcase, GraduationCap, ArrowUpRight, Sparkles } from 'lucide-svelte';

  let list = $state<Alumni[]>([]);

  onMount(async () => {
    try {
      list = await api.getAlumni();
    } catch (e) {
      console.error(e);
    }
  });
</script>

<svelte:head>
  <title>Ikatan Alumni Al-Hamid (IKAH) | Pondok Pesantren Al-Hamid</title>
</svelte:head>

<!-- Header -->
<div class="bg-[#090D16] text-white py-20 px-4 sm:px-8 relative overflow-hidden border-b border-slate-800/80">
  <div class="absolute inset-0 bg-grid-dark opacity-30 pointer-events-none"></div>
  <div class="absolute -top-24 left-1/2 -translate-x-1/2 w-[600px] h-[300px] bg-teal-500/15 blur-3xl pointer-events-none rounded-full"></div>

  <div class="max-w-6xl mx-auto space-y-4 relative z-10">
    <div class="inline-flex items-center gap-2 px-3.5 py-1.5 rounded-full bg-teal-500/15 text-teal-300 text-xs font-bold border border-teal-500/30">
      <Sparkles class="w-3.5 h-3.5 text-teal-400" />
      <span>Jejaring & Pengabdian</span>
    </div>
    <h1 class="font-display font-black text-3xl sm:text-5xl lg:text-6xl text-white tracking-tight">
      Ikatan Alumni Pondok Pesantren Al-Hamid
    </h1>
    <p class="text-slate-300 text-base sm:text-lg max-w-2xl font-normal leading-relaxed">
      Wadah sinergi ribuan alumni Al-Hamid yang berkiprah di sektor keagamaan, akademisi, sains teknologi, birokrasi, dan industri kreatif di berbagai belahan dunia.
    </p>
  </div>
</div>

<div class="max-w-6xl mx-auto px-4 sm:px-8 py-20">
  <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-8">
    {#each list as item}
      <div class="bg-white rounded-3xl border border-slate-200/80 p-6 sm:p-7 shadow-sm hover:shadow-2xl hover:-translate-y-1 transition-all duration-300 flex flex-col justify-between space-y-4 group">
        <div class="space-y-4">
          <div class="flex items-center gap-4">
            <img
              src={item.photo_url}
              alt={item.name}
              class="w-16 h-16 rounded-2xl object-cover border-2 border-teal-500 shadow-md group-hover:scale-105 transition"
            />
            <div>
              <h2 class="font-display font-bold text-base text-slate-950 group-hover:text-teal-700 transition leading-snug">{item.name}</h2>
              <div class="text-xs text-teal-700 font-black uppercase tracking-wider mt-0.5">Angkatan {item.graduation_year}</div>
              <div class="text-[11px] text-slate-400">{item.education_unit}</div>
            </div>
          </div>

          <div class="space-y-2 text-xs text-slate-600 pt-1">
            <div class="flex items-center gap-2">
              <Briefcase class="w-3.5 h-3.5 text-teal-600 shrink-0" />
              <span class="font-medium text-slate-700">{item.profession}</span>
            </div>
            <div class="flex items-center gap-2">
              <MapPin class="w-3.5 h-3.5 text-teal-600 shrink-0" />
              <span>{item.city}</span>
            </div>
          </div>

          {#if item.testimonial}
            <p class="font-editorial italic text-xs text-slate-600 pt-2 border-t border-slate-100 leading-relaxed">
              "{item.testimonial}"
            </p>
          {/if}
        </div>
      </div>
    {/each}
  </div>
</div>
