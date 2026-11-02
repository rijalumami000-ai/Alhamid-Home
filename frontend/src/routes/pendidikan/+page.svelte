<script lang="ts">
  import { onMount } from 'svelte';
  import { api } from '$lib/api';
  import type { EducationUnit } from '$lib/types';
  import { GraduationCap, BookOpen, ArrowRight, CheckCircle2 } from 'lucide-svelte';

  let units = $state<EducationUnit[]>([]);
  let activeTab = $state<'all' | 'formal' | 'non_formal'>('all');

  onMount(async () => {
    try {
      units = await api.getEducationUnits();
    } catch (e) {
      console.error(e);
    }
  });

  let filteredUnits = $derived(
    activeTab === 'all' ? units : units.filter((u) => u.type === activeTab)
  );
</script>

<svelte:head>
  <title>Pendidikan Formal & Non-Formal | Pondok Pesantren Al-Hamid</title>
</svelte:head>

<div class="bg-gradient-to-r from-emerald-950 via-emerald-900 to-emerald-800 text-white py-16 px-4 sm:px-8">
  <div class="max-w-6xl mx-auto space-y-4">
    <div class="inline-flex items-center gap-2 px-3.5 py-1.5 rounded-full bg-emerald-500/20 text-emerald-300 text-xs font-semibold border border-emerald-500/30">
      Program Studi & Jenjang
    </div>
    <h1 class="font-serif-luxury font-bold text-3xl sm:text-5xl text-white">
      Pendidikan Pondok Pesantren Al-Hamid
    </h1>
    <p class="text-emerald-100/80 text-base max-w-2xl">
      Paduan terintegrasi antara pendidikan non-formal (Madrasah Diniyah & PMQ) serta pendidikan formal (MTs, MA, dan MAI Cintamulya).
    </p>
  </div>
</div>

<div class="max-w-6xl mx-auto px-4 sm:px-8 py-16 space-y-10">
  <!-- Tabs -->
  <div class="flex items-center gap-2 border-b border-slate-200 pb-4">
    <button
      onclick={() => (activeTab = 'all')}
      class="px-5 py-2.5 rounded-xl text-xs sm:text-sm font-semibold transition {activeTab === 'all'
        ? 'bg-emerald-800 text-white shadow-md'
        : 'text-slate-600 hover:bg-slate-100'}"
    >
      Semua Jenjang
    </button>
    <button
      onclick={() => (activeTab = 'formal')}
      class="px-5 py-2.5 rounded-xl text-xs sm:text-sm font-semibold transition {activeTab === 'formal'
        ? 'bg-emerald-800 text-white shadow-md'
        : 'text-slate-600 hover:bg-slate-100'}"
    >
      Pendidikan Formal (MTs, MA, MAI)
    </button>
    <button
      onclick={() => (activeTab = 'non_formal')}
      class="px-5 py-2.5 rounded-xl text-xs sm:text-sm font-semibold transition {activeTab === 'non_formal'
        ? 'bg-emerald-800 text-white shadow-md'
        : 'text-slate-600 hover:bg-slate-100'}"
    >
      Pendidikan Non-Formal (Diniyah & PMQ)
    </button>
  </div>

  <!-- Units Grid -->
  <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-8">
    {#each filteredUnits as unit}
      <div class="bg-white rounded-3xl border border-slate-100 shadow-md hover:shadow-2xl transition-all duration-300 overflow-hidden flex flex-col group">
        <div class="relative h-52 overflow-hidden bg-slate-900">
          <img
            src={unit.hero_image}
            alt={unit.name}
            class="w-full h-full object-cover group-hover:scale-105 transition duration-500"
          />
          <div class="absolute inset-0 bg-gradient-to-t from-slate-950/80 via-transparent to-transparent"></div>
          <div class="absolute top-3 right-3">
            <span class="px-2.5 py-1 rounded-full text-[11px] font-bold uppercase shadow-sm {unit.type === 'formal' ? 'bg-blue-600 text-white' : 'bg-emerald-700 text-white'}">
              {unit.type === 'formal' ? 'Formal' : 'Non-Formal'}
            </span>
          </div>
          <div class="absolute bottom-3 left-4 right-4">
            <span class="text-amber-300 text-xs font-semibold">{unit.badge}</span>
          </div>
        </div>

        <div class="p-6 flex-1 flex flex-col justify-between space-y-4">
          <div>
            <h2 class="font-serif-luxury font-bold text-xl text-slate-900 group-hover:text-emerald-700 transition">
              {unit.name}
            </h2>
            <p class="text-xs font-medium text-slate-400 mt-1">{unit.tagline}</p>
            <p class="text-slate-600 text-xs sm:text-sm mt-3 line-clamp-3 leading-relaxed">
              {unit.short_desc}
            </p>
          </div>

          <div class="pt-3 border-t border-slate-100 flex items-center justify-between">
            <span class="text-xs text-slate-400">Kepala: {unit.headmaster_name}</span>
            <a
              href="/pendidikan/{unit.slug}"
              class="inline-flex items-center gap-1.5 px-4 py-2 rounded-xl bg-emerald-50 hover:bg-emerald-100 text-emerald-800 font-bold text-xs transition"
            >
              <span>Laman Lengkap</span>
              <ArrowRight class="w-3.5 h-3.5" />
            </a>
          </div>
        </div>
      </div>
    {/each}
  </div>
</div>

