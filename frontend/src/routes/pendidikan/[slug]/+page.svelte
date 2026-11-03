<script lang="ts">
  import { onMount } from 'svelte';
  import { page } from '$app/state';
  import { api } from '$lib/api';
  import type { EducationUnit } from '$lib/types';
  import {
    GraduationCap,
    BookOpen,
    CheckCircle2,
    Calendar,
    Users,
    Building2,
    ClipboardList,
    ArrowRight,
    Sparkles,
    ShieldCheck
  } from 'lucide-svelte';

  let unit = $state<EducationUnit | null>(null);
  let loading = $state(true);
  let errorMsg = $state<string | null>(null);

  async function loadUnit(slug: string) {
    loading = true;
    errorMsg = null;
    try {
      unit = await api.getEducationUnitBySlug(slug);
    } catch (e: any) {
      errorMsg = e.message || 'Laman unit pendidikan tidak ditemukan';
    } finally {
      loading = false;
    }
  }

  $effect(() => {
    const slug = page.params.slug;
    if (slug) {
      loadUnit(slug);
    }
  });
</script>

<svelte:head>
  <title>{unit?.name ? `${unit.name} | Pondok Pesantren Al-Hamid` : 'Laman Pendidikan Al-Hamid'}</title>
</svelte:head>

{#if loading}
  <div class="py-32 text-center text-slate-400">
    <div class="animate-spin w-8 h-8 border-4 border-teal-600 border-t-transparent rounded-full mx-auto mb-4"></div>
    Memuat informasi unit pendidikan...
  </div>
{:else if errorMsg || !unit}
  <div class="py-32 text-center max-w-md mx-auto px-4 space-y-4">
    <h1 class="text-2xl font-display font-black text-slate-800">Unit Pendidikan Tidak Ditemukan</h1>
    <p class="text-sm text-slate-500">Laman yang Anda tuju mungkin telah dipindahkan atau belum tersedia.</p>
    <a href="/pendidikan" class="inline-flex px-6 py-2.5 rounded-full bg-slate-900 text-white font-bold text-xs uppercase tracking-wider">
      Kembali ke Daftar Pendidikan
    </a>
  </div>
{:else}
  <!-- Hero Section -->
  <section class="relative min-h-[460px] flex items-center bg-[#090D16] text-white overflow-hidden border-b border-slate-800/80">
    <div class="absolute inset-0">
      <img
        src={unit.hero_image}
        alt={unit.name}
        class="w-full h-full object-cover opacity-35 mix-blend-luminosity"
      />
      <div class="absolute inset-0 bg-gradient-to-r from-[#090D16] via-[#090D16]/85 to-transparent"></div>
      <div class="absolute inset-0 bg-grid-dark opacity-30 pointer-events-none"></div>
    </div>

    <div class="relative z-10 max-w-6xl mx-auto px-4 sm:px-8 py-20 w-full space-y-4">
      <div class="flex flex-wrap items-center gap-3">
        <span class="px-3.5 py-1 rounded-full text-[11px] font-black uppercase tracking-wider shadow-sm {unit.type === 'formal' ? 'bg-indigo-600 text-white' : 'bg-teal-600 text-white'}">
          {unit.type === 'formal' ? 'Pendidikan Formal' : 'Pendidikan Non-Formal'}
        </span>
        {#if unit.badge}
          <span class="px-3.5 py-1 rounded-full bg-amber-400/15 text-amber-300 border border-amber-400/30 text-xs font-bold backdrop-blur-md">
            {unit.badge}
          </span>
        {/if}
      </div>

      <h1 class="font-display font-black text-3xl sm:text-5xl lg:text-6xl text-white tracking-tight">
        {unit.name}
      </h1>

      <p class="text-base sm:text-lg text-slate-300 max-w-2xl font-normal leading-relaxed">
        {unit.tagline}
      </p>

      <div class="pt-4 flex items-center gap-4">
        <a
          href="/kontak"
          class="px-8 py-3.5 rounded-full bg-gradient-to-r from-teal-400 to-emerald-400 hover:from-teal-300 hover:to-emerald-300 text-slate-950 font-display font-black text-xs uppercase tracking-wider shadow-lg hover:scale-105 active:scale-95 transition"
        >
          Daftar ke {unit.name}
        </a>
      </div>
    </div>
  </section>

  <!-- Content Body -->
  <div class="max-w-6xl mx-auto px-4 sm:px-8 py-20">
    <div class="grid grid-cols-1 lg:grid-cols-12 gap-12">
      <!-- Main Column -->
      <div class="lg:col-span-8 space-y-12">
        <!-- Gambaran Umum -->
        <div class="space-y-4">
          <div class="text-xs font-black uppercase tracking-widest text-teal-700 flex items-center gap-1.5">
            <BookOpen class="w-4 h-4" /> Gambaran Umum Program
          </div>
          <h2 class="font-display font-black text-2xl sm:text-3xl text-slate-950 tracking-tight">
            Membangun Karakter & Kapasitas Intelektual
          </h2>
          <div class="text-slate-600 leading-relaxed text-base whitespace-pre-line bg-white p-8 sm:p-10 rounded-3xl border border-slate-200/80 shadow-sm">
            {unit.full_desc}
          </div>
        </div>

        <!-- Kurikulum & Kitab / Pembelajaran -->
        {#if unit.curriculum}
          <div class="space-y-4">
            <div class="text-xs font-black uppercase tracking-widest text-teal-700 flex items-center gap-1.5">
              <ClipboardList class="w-4 h-4" /> Kurikulum & Silabus Pembelajaran
            </div>
            <h2 class="font-display font-black text-2xl sm:text-3xl text-slate-950 tracking-tight">
              Struktur Kurikulum Terpadu
            </h2>
            <div class="bg-[#090D16] text-slate-200 p-8 sm:p-10 rounded-3xl border border-slate-800 shadow-xl space-y-4 text-sm sm:text-base leading-relaxed whitespace-pre-line relative overflow-hidden">
              <div class="absolute inset-0 bg-grid-dark opacity-20 pointer-events-none"></div>
              <div class="relative z-10">{unit.curriculum}</div>
            </div>
          </div>
        {/if}

        <!-- Persyaratan Masuk -->
        {#if unit.requirements}
          <div class="space-y-4">
            <div class="text-xs font-black uppercase tracking-widest text-teal-700 flex items-center gap-1.5">
              <CheckCircle2 class="w-4 h-4" /> Persyaratan & Alur Masuk
            </div>
            <h2 class="font-display font-black text-2xl sm:text-3xl text-slate-950 tracking-tight">
              Kriteria Penerimaan Santri Baru
            </h2>
            <div class="bg-white p-8 sm:p-10 rounded-3xl border border-slate-200/80 shadow-sm text-slate-600 text-sm sm:text-base leading-relaxed whitespace-pre-line">
              {unit.requirements}
            </div>
          </div>
        {/if}
      </div>

      <!-- Sidebar -->
      <div class="lg:col-span-4 space-y-6">
        <!-- Kepala Madrasah / Pimpinan Unit -->
        <div class="bg-white rounded-3xl border border-slate-200/80 p-7 shadow-sm space-y-4">
          <div class="text-xs font-black uppercase tracking-widest text-teal-700">Pimpinan Lembaga</div>
          <div class="flex items-center gap-4">
            <div class="w-14 h-14 rounded-2xl bg-teal-50 text-teal-700 flex items-center justify-center font-display font-black text-xl shrink-0">
              {unit.headmaster_name.charAt(0)}
            </div>
            <div>
              <h3 class="font-display font-bold text-base text-slate-950 leading-snug">{unit.headmaster_name}</h3>
              <div class="text-xs text-slate-500 mt-0.5">Kepala / Mudir Unit</div>
            </div>
          </div>
        </div>

        <!-- Quick Contacts & PSB CTA Card -->
        <div class="bg-gradient-to-br from-[#090D16] to-slate-950 text-white rounded-3xl p-7 border border-slate-800 space-y-4 relative overflow-hidden">
          <div class="inline-flex items-center gap-1.5 px-3 py-1 rounded-full bg-teal-500/20 text-teal-300 text-xs font-bold">
            <Sparkles class="w-3.5 h-3.5 text-teal-400" />
            <span>Pendaftaran 2026/2027</span>
          </div>
          <h3 class="font-display font-black text-xl text-white tracking-tight">Konsultasi Pendaftaran</h3>
          <p class="text-xs text-slate-400 leading-relaxed font-normal">
            Dapatkan informasi lengkap mengenai kurikulum, asrama, dan beasiswa pendidikan {unit.name}.
          </p>
          <div class="pt-2 flex flex-col gap-2.5">
            <a
              href="/kontak"
              class="w-full text-center py-3 rounded-full bg-teal-500 hover:bg-teal-400 text-slate-950 font-display font-bold text-xs uppercase tracking-wider shadow-md transition"
            >
              Formulir Pendaftaran Online
            </a>
            <a
              href="https://wa.me/6281234567890"
              target="_blank"
              rel="noopener noreferrer"
              class="w-full text-center py-3 rounded-full bg-white/10 hover:bg-white/15 text-white font-display font-bold text-xs tracking-wider border border-white/20 transition"
            >
              Chat WhatsApp Sekretariat
            </a>
          </div>
        </div>
      </div>
    </div>
  </div>
{/if}
