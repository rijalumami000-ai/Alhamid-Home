<script lang="ts">
  import { onMount } from 'svelte';
  import { api } from '$lib/api';
  import type { HomeData } from '$lib/types';
  import {
    ArrowRight,
    Sparkles,
    GraduationCap,
    BookOpen,
    Trophy,
    Building2,
    Users,
    CheckCircle2,
    Play,
    Star,
    Calendar,
    ChevronRight,
    Compass,
    ArrowUpRight,
    Shield,
    ExternalLink,
    Quote
  } from 'lucide-svelte';

  let data = $state<HomeData | null>(null);
  let loading = $state(true);
  let activeBannerIndex = $state(0);
  let videoModalUrl = $state<string | null>(null);

  onMount(async () => {
    try {
      data = await api.getHome();
    } catch (e) {
      console.error('Failed to load home data:', e);
    } finally {
      loading = false;
    }

    // Auto rotate banners if more than 1
    const interval = setInterval(() => {
      if (data && data.banners.length > 1) {
        activeBannerIndex = (activeBannerIndex + 1) % data.banners.length;
      }
    }, 7000);

    return () => clearInterval(interval);
  });
</script>

<svelte:head>
  <title>Pondok Pesantren Al-Hamid | Berakhlak, Unggul & Berdaya Saing Global</title>
  <meta name="description" content="Pondok Pesantren Al-Hamid Cintamulya: Paduan harmoni antara tradisi luhur kitab kuning dan keunggulan sains modern (MTs, MA, MAI Kejuruan IT, Diniyah, PMQ)." />
</svelte:head>

<!-- ================= 1. HERO SECTION: CINEMATIC & ARCHITECTURAL ================= -->
<section class="relative min-h-[680px] lg:min-h-[760px] flex items-center bg-[#090D16] text-white overflow-hidden">
  <!-- Architectural ambient glow and grid -->
  <div class="absolute inset-0 bg-grid-dark opacity-35 pointer-events-none"></div>
  <div class="absolute -top-32 left-1/2 -translate-x-1/2 w-[800px] h-[400px] bg-gradient-to-b from-teal-500/20 via-cyan-500/10 to-transparent blur-3xl pointer-events-none rounded-full"></div>
  <div class="absolute -bottom-24 right-0 w-[500px] h-[300px] bg-amber-500/10 blur-3xl pointer-events-none rounded-full"></div>

  {#if data && data.banners.length > 0}
    {#each data.banners as banner, i}
      <div
        class="absolute inset-0 transition-opacity duration-1000 ease-in-out {i === activeBannerIndex ? 'opacity-100 z-10' : 'opacity-0 z-0'}"
      >
        <!-- Background Image with Depth Gradient -->
        <img
          src={banner.image_url}
          alt={banner.title}
          class="w-full h-full object-cover object-center scale-105 transform transition-transform duration-10000 opacity-40 mix-blend-luminosity"
        />
        <div class="absolute inset-0 bg-gradient-to-r from-[#090D16] via-[#090D16]/80 to-[#090D16]/30"></div>
        <div class="absolute inset-0 bg-gradient-to-t from-[#090D16] via-transparent to-[#090D16]/70"></div>
      </div>
    {/each}

    <!-- Hero Content Overlay -->
    <div class="relative z-20 max-w-7xl mx-auto px-4 sm:px-8 py-24 sm:py-32 w-full">
      <div class="max-w-3xl space-y-6">
        <!-- Interactive Pill Tag -->
        {#if data.banners[activeBannerIndex]?.badge}
          <div class="inline-flex items-center gap-2.5 px-4 py-1.5 rounded-full bg-teal-500/15 border border-teal-400/30 text-teal-300 text-xs font-semibold backdrop-blur-xl shadow-lg shadow-teal-950/20 animate-in fade-in duration-300">
            <span class="w-2 h-2 rounded-full bg-teal-400 animate-ping"></span>
            <span>{data.banners[activeBannerIndex].badge}</span>
          </div>
        {/if}

        <!-- Master Headline -->
        <h1 class="font-display font-black text-4xl sm:text-6xl lg:text-7xl text-white tracking-tight leading-[1.08]">
          {data.banners[activeBannerIndex]?.title}
        </h1>

        <!-- Subtitle -->
        <p class="text-base sm:text-xl text-slate-300 leading-relaxed max-w-2xl font-normal">
          {data.banners[activeBannerIndex]?.subtitle}
        </p>

        <!-- CTA Action Buttons -->
        <div class="pt-4 flex flex-wrap items-center gap-4">
          <a
            href={data.banners[activeBannerIndex]?.button_link || '/kontak'}
            class="px-8 py-4 rounded-full bg-gradient-to-r from-teal-500 via-emerald-500 to-teal-600 hover:from-teal-400 hover:to-emerald-400 text-slate-950 font-display font-extrabold text-sm tracking-wide shadow-xl shadow-teal-900/30 hover:scale-105 active:scale-95 transition duration-200 flex items-center gap-2.5"
          >
            <span>{data.banners[activeBannerIndex]?.button_text || 'Jelajahi Sekarang'}</span>
            <ArrowRight class="w-4 h-4" />
          </a>
          <a
            href="/pendidikan"
            class="px-8 py-4 rounded-full bg-white/10 hover:bg-white/15 text-white font-display font-bold text-sm tracking-wide backdrop-blur-xl border border-white/20 hover:border-white/40 transition duration-200 flex items-center gap-2"
          >
            <span>Program Akademik</span>
            <Compass class="w-4 h-4 text-teal-400" />
          </a>
        </div>
      </div>

      <!-- Slide Segment Indicators -->
      <div class="absolute bottom-8 left-4 sm:left-8 flex items-center gap-2.5 z-20">
        {#each data.banners as _, i}
          <button
            onclick={() => (activeBannerIndex = i)}
            class="h-2 rounded-full transition-all duration-300 {i === activeBannerIndex ? 'w-10 bg-teal-400' : 'w-2.5 bg-white/30 hover:bg-white/50'}"
            aria-label="Slide {i + 1}"
          ></button>
        {/each}
      </div>
    </div>
  {:else}
    <div class="relative z-20 max-w-7xl mx-auto px-4 sm:px-8 py-32 text-center">
      <div class="inline-flex items-center gap-2 px-4 py-1.5 rounded-full bg-teal-500/20 text-teal-300 text-xs font-semibold mb-4">
        Pondok Pesantren Al-Hamid
      </div>
      <h1 class="font-display font-black text-4xl sm:text-6xl lg:text-7xl text-white mb-6 tracking-tight">
        Menempa Akhlak, Mengukir Prestasi Masa Depan
      </h1>
      <p class="text-slate-300 max-w-2xl mx-auto mb-8 text-lg">
        Paduan harmoni antara kedalaman turats pesantren salaf dan keunggulan sains modern.
      </p>
    </div>
  {/if}
</section>

<!-- ================= 2. QUICK STATS BENTO STRIP ================= -->
<section class="relative z-30 -mt-12 max-w-7xl mx-auto px-4 sm:px-8">
  <div class="bg-white/95 backdrop-blur-xl rounded-3xl shadow-2xl shadow-slate-950/5 border border-slate-200/80 p-6 sm:p-8 grid grid-cols-2 lg:grid-cols-4 gap-6 sm:gap-8">
    <div class="space-y-1 sm:space-y-2">
      <div class="flex items-center gap-2 text-teal-600 text-xs font-bold uppercase tracking-wider">
        <Users class="w-4 h-4" />
        <span>Santri Aktif</span>
      </div>
      <div class="font-display font-black text-3xl sm:text-4xl text-slate-900 tracking-tight">
        {data?.stats?.santri_count || '1.450+'}
      </div>
      <div class="text-xs text-slate-500 font-medium">Santri Mukim Terbina Intensif</div>
    </div>

    <div class="space-y-1 sm:space-y-2 border-l border-slate-100 pl-4 sm:pl-8">
      <div class="flex items-center gap-2 text-indigo-600 text-xs font-bold uppercase tracking-wider">
        <GraduationCap class="w-4 h-4" />
        <span>Dewan Asatidz</span>
      </div>
      <div class="font-display font-black text-3xl sm:text-4xl text-slate-900 tracking-tight">
        {data?.stats?.asatidz_count || '85+'}
      </div>
      <div class="text-xs text-slate-500 font-medium">Masyayikh & Pendidik Berkualitas</div>
    </div>

    <div class="space-y-1 sm:space-y-2 border-l-0 lg:border-l border-slate-100 pl-0 lg:pl-8">
      <div class="flex items-center gap-2 text-amber-600 text-xs font-bold uppercase tracking-wider">
        <Trophy class="w-4 h-4" />
        <span>Prestasi Santri</span>
      </div>
      <div class="font-display font-black text-3xl sm:text-4xl text-slate-900 tracking-tight">
        {data?.stats?.prestasi_count || '120+'}
      </div>
      <div class="text-xs text-slate-500 font-medium">Juara Nasional & Internasional</div>
    </div>

    <div class="space-y-1 sm:space-y-2 border-l border-slate-100 pl-4 sm:pl-8">
      <div class="flex items-center gap-2 text-emerald-600 text-xs font-bold uppercase tracking-wider">
        <Compass class="w-4 h-4" />
        <span>Jejaring Alumni</span>
      </div>
      <div class="font-display font-black text-3xl sm:text-4xl text-slate-900 tracking-tight">
        {data?.stats?.alumni_count || '6.200+'}
      </div>
      <div class="text-xs text-slate-500 font-medium">Kiprah di 34 Provinsi & Luar Negeri</div>
    </div>
  </div>
</section>

<!-- ================= 3. EDITORIAL NARRATIVE & SAMBUTAN PENGASUH ================= -->
<section class="py-24 max-w-7xl mx-auto px-4 sm:px-8">
  <div class="grid grid-cols-1 lg:grid-cols-12 gap-12 lg:gap-16 items-center">
    <!-- Editorial Photography Frame -->
    <div class="lg:col-span-5 relative">
      <div class="relative rounded-3xl overflow-hidden shadow-2xl border-8 border-white bg-slate-950 aspect-[4/5] group">
        <img
          src="https://images.unsplash.com/photo-1542838132-92c53300491e?auto=format&fit=crop&w=1000&q=80"
          alt="Kajian Kitab Pondok Pesantren Al-Hamid"
          class="w-full h-full object-cover group-hover:scale-105 transition duration-700 opacity-90"
        />
        <div class="absolute inset-0 bg-gradient-to-t from-slate-950 via-slate-950/40 to-transparent"></div>
        <div class="absolute bottom-6 left-6 right-6 text-white space-y-1">
          <div class="inline-flex items-center gap-2 px-3 py-1 rounded-full bg-teal-500/20 text-teal-300 text-[11px] font-bold uppercase tracking-wider backdrop-blur-md">
            Mursyid & Pengasuh Utama
          </div>
          <div class="font-display font-extrabold text-2xl text-white">
            K.H. Ahmad Hamidurrohman
          </div>
          <p class="font-editorial italic text-slate-300 text-sm">
            "Menjaga kemurnian sanad, merawat tradisi turats, dan merajut keunggulan sains peradaban."
          </p>
        </div>
      </div>

      <!-- Floating Credential Card -->
      <div class="absolute -bottom-6 -right-6 bg-slate-900 text-white p-5 rounded-2xl shadow-2xl border border-slate-800 flex items-center gap-3.5 hidden sm:flex max-w-xs">
        <div class="w-10 h-10 rounded-xl bg-teal-500/20 text-teal-400 flex items-center justify-center shrink-0">
          <Shield class="w-5 h-5" />
        </div>
        <div>
          <div class="font-display font-bold text-xs text-white">Sanad Keilmuan Muttashil</div>
          <div class="text-[11px] text-slate-400 mt-0.5">Jalur Masyayikh Lirboyo & Sarang</div>
        </div>
      </div>
    </div>

    <!-- Narrative & Values -->
    <div class="lg:col-span-7 space-y-6">
      <div class="inline-flex items-center gap-2 px-3.5 py-1.5 rounded-full bg-teal-50 border border-teal-200/60 text-teal-800 text-xs font-bold tracking-wide">
        <Sparkles class="w-3.5 h-3.5 text-teal-600" />
        <span>Visi Pendidikan Komprehensif</span>
      </div>

      <h2 class="font-display font-black text-3xl sm:text-5xl text-slate-950 leading-[1.15] tracking-tight">
        Mencetak Generasi Intelektual Beradab & Berkecakapan Global
      </h2>

      <p class="text-slate-600 text-base sm:text-lg leading-relaxed font-normal">
        Pondok Pesantren Al-Hamid memosisikan diri sebagai inkubator kader pemimpin umat yang tangguh. Kami memadukan kedalaman kajian teks klasik (turats) dengan penguasaan sains modern, teknologi komputasi, dan kepemimpinan berwawasan kebangsaan.
      </p>

      <div class="grid grid-cols-1 sm:grid-cols-2 gap-4 pt-2">
        <div class="p-5 rounded-2xl bg-white border border-slate-200/70 shadow-sm hover:border-teal-500/40 transition space-y-2">
          <div class="w-10 h-10 rounded-xl bg-teal-50 text-teal-700 flex items-center justify-center">
            <BookOpen class="w-5 h-5" />
          </div>
          <div class="font-display font-bold text-base text-slate-900">Salafiyah & Sanad Talaqqi</div>
          <p class="text-xs text-slate-500 leading-relaxed">
            Pengajian kitab kuning terstruktur dari dasar hingga tingkat tinggi, serta tahfidz bersanad Al-Qur'an.
          </p>
        </div>

        <div class="p-5 rounded-2xl bg-white border border-slate-200/70 shadow-sm hover:border-indigo-500/40 transition space-y-2">
          <div class="w-10 h-10 rounded-xl bg-indigo-50 text-indigo-700 flex items-center justify-center">
            <GraduationCap class="w-5 h-5" />
          </div>
          <div class="font-display font-bold text-base text-slate-900">Sains & Kejuruan Digital</div>
          <p class="text-xs text-slate-500 leading-relaxed">
            MTs & MA Cintamulya (Akreditasi A), serta MAI Kejuruan Informatika dengan kurikulum teknologi industri.
          </p>
        </div>
      </div>

      <div class="pt-2 flex items-center gap-4">
        <a
          href="/tentang/profil"
          class="inline-flex items-center gap-2 text-sm font-display font-bold text-teal-700 hover:text-teal-800 group"
        >
          <span>Eksplorasi Profil & Sejarah Selengkapnya</span>
          <ArrowRight class="w-4 h-4 group-hover:translate-x-1.5 transition" />
        </a>
      </div>
    </div>
  </div>
</section>

<!-- ================= 4. PROGRAM PENDIDIKAN BENTO CARDS ================= -->
<section class="py-24 bg-slate-100/70 border-y border-slate-200/70 relative">
  <div class="max-w-7xl mx-auto px-4 sm:px-8">
    <div class="flex flex-col md:flex-row md:items-end justify-between gap-6 mb-16">
      <div class="space-y-3">
        <div class="inline-flex items-center gap-2 px-3.5 py-1.5 rounded-full bg-teal-50 border border-teal-200/70 text-teal-800 text-xs font-bold">
          Ekosistem Pendidikan
        </div>
        <h2 class="font-display font-black text-3xl sm:text-5xl text-slate-950 tracking-tight">
          Jenjang Formal & Non-Formal Terpadu
        </h2>
      </div>
      <p class="text-slate-600 text-sm sm:text-base max-w-md">
        Kurikulum yang saling melengkapi untuk memastikan santri matang dalam spiritualitas sekaligus kompetitif secara profesional.
      </p>
    </div>

    <!-- Asymmetric Bento Cards -->
    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6 sm:gap-8">
      {#if data?.education_units}
        {#each data.education_units as unit}
          <div class="bg-white rounded-3xl border border-slate-200/80 shadow-md hover:shadow-2xl hover:-translate-y-1.5 transition-all duration-300 overflow-hidden flex flex-col group">
            <!-- Header Image -->
            <div class="relative h-52 overflow-hidden bg-slate-950">
              <img
                src={unit.hero_image}
                alt={unit.name}
                class="w-full h-full object-cover group-hover:scale-105 transition duration-700"
              />
              <div class="absolute inset-0 bg-gradient-to-t from-slate-950/85 via-slate-950/20 to-transparent"></div>
              
              <div class="absolute top-4 left-4 right-4 flex items-center justify-between">
                <span class="px-3 py-1 rounded-full text-[10px] font-black tracking-wider uppercase shadow-md {unit.type === 'formal' ? 'bg-indigo-600 text-white' : 'bg-teal-600 text-white'}">
                  {unit.type === 'formal' ? 'Pendidikan Formal' : 'Pendidikan Non-Formal'}
                </span>
                <span class="text-amber-300 text-xs font-bold drop-shadow">{unit.badge}</span>
              </div>

              <div class="absolute bottom-3 left-4 right-4">
                <span class="text-white text-xs font-medium opacity-90">{unit.tagline}</span>
              </div>
            </div>

            <!-- Content -->
            <div class="p-6 flex-1 flex flex-col justify-between space-y-4">
              <div>
                <h3 class="font-display font-black text-xl text-slate-950 group-hover:text-teal-700 transition">
                  {unit.name}
                </h3>
                <p class="text-slate-600 text-xs sm:text-sm mt-3 line-clamp-3 leading-relaxed">
                  {unit.short_desc}
                </p>
              </div>

              <div class="pt-4 border-t border-slate-100 flex items-center justify-between">
                <span class="text-[11px] font-semibold text-slate-400">Kepala: {unit.headmaster_name}</span>
                <a
                  href="/pendidikan/{unit.slug}"
                  class="inline-flex items-center gap-1.5 text-xs font-display font-bold text-teal-700 hover:text-teal-800 group/link"
                >
                  <span>Laman Unit</span>
                  <ArrowRight class="w-3.5 h-3.5 group-hover/link:translate-x-1 transition" />
                </a>
              </div>
            </div>
          </div>
        {/each}
      {/if}
    </div>
  </div>
</section>

<!-- ================= 5. HIGHLIGHT BERITA & KAJIAN ================= -->
<section class="py-24 max-w-7xl mx-auto px-4 sm:px-8">
  <div class="flex flex-col md:flex-row md:items-end justify-between gap-6 mb-16">
    <div class="space-y-3">
      <div class="inline-flex items-center gap-2 px-3.5 py-1.5 rounded-full bg-teal-50 border border-teal-200/70 text-teal-800 text-xs font-bold">
        Warta & Literasi
      </div>
      <h2 class="font-display font-black text-3xl sm:text-5xl text-slate-950 tracking-tight">
        Kajian, Opini & Dinamika Santri
      </h2>
    </div>
    <a
      href="/berita"
      class="inline-flex items-center gap-2 text-sm font-display font-bold text-teal-700 hover:text-teal-800 group"
    >
      <span>Lihat Arsip Berita Selengkapnya</span>
      <ArrowRight class="w-4 h-4 group-hover:translate-x-1.5 transition" />
    </a>
  </div>

  <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-8">
    {#if data?.featured_articles}
      {#each data.featured_articles as article}
        <article class="bg-white rounded-3xl border border-slate-200/80 shadow-sm hover:shadow-xl hover:-translate-y-1 transition-all duration-300 overflow-hidden flex flex-col group">
          <a href="/berita/{article.slug}" class="relative h-52 overflow-hidden bg-slate-950 block">
            <img
              src={article.thumbnail_url}
              alt={article.title}
              class="w-full h-full object-cover group-hover:scale-105 transition duration-700"
            />
            <div class="absolute inset-0 bg-gradient-to-t from-slate-950/70 via-transparent to-transparent"></div>
            <div class="absolute top-3 left-3">
              <span class="px-3 py-1 rounded-full bg-slate-900/90 text-teal-300 text-[11px] font-bold backdrop-blur-md border border-white/10">
                {article.category?.name || 'Kajian'}
              </span>
            </div>
          </a>

          <div class="p-6 flex-1 flex flex-col justify-between space-y-4">
            <div>
              <div class="text-xs text-slate-400 mb-2 flex items-center gap-2 font-medium">
                <Calendar class="w-3.5 h-3.5 text-teal-600" />
                <span>{new Date(article.published_at).toLocaleDateString('id-ID', { day: 'numeric', month: 'long', year: 'numeric' })}</span>
                <span>&bull;</span>
                <span>{article.author}</span>
              </div>
              <h3 class="font-display font-bold text-lg text-slate-950 group-hover:text-teal-700 transition line-clamp-2 leading-snug">
                <a href="/berita/{article.slug}">{article.title}</a>
              </h3>
              <p class="text-slate-600 text-xs sm:text-sm mt-2 line-clamp-2 leading-relaxed">
                {article.excerpt}
              </p>
            </div>

            <div class="pt-4 border-t border-slate-100 flex items-center justify-between">
              <span class="text-xs text-slate-400 font-medium">{article.view_count} pembaca</span>
              <a
                href="/berita/{article.slug}"
                class="inline-flex items-center gap-1 text-xs font-display font-bold text-teal-700 hover:text-teal-800"
              >
                <span>Baca Lengkap</span>
                <ChevronRight class="w-3.5 h-3.5" />
              </a>
            </div>
          </div>
        </article>
      {/each}
    {/if}
  </div>
</section>

<!-- ================= 6. OBSIDIAN SHOWCASE: PRESTASI & FASILITAS ================= -->
<section class="py-24 bg-[#090D16] text-white border-y border-slate-800/80 relative overflow-hidden">
  <div class="absolute inset-0 bg-grid-dark opacity-20 pointer-events-none"></div>

  <div class="max-w-7xl mx-auto px-4 sm:px-8 relative z-10">
    <div class="grid grid-cols-1 lg:grid-cols-2 gap-16 items-start">
      <!-- Achievements Showcase -->
      <div class="space-y-6">
        <div class="inline-flex items-center gap-2 px-3.5 py-1.5 rounded-full bg-amber-500/15 text-amber-300 text-xs font-bold border border-amber-500/30">
          <Trophy class="w-3.5 h-3.5 text-amber-400" />
          <span>Prestasi Gemilang</span>
        </div>

        <h2 class="font-display font-black text-3xl sm:text-5xl text-white tracking-tight">
          Rekam Jejak Juara di Panggung Nasional
        </h2>

        <p class="text-slate-300 text-sm sm:text-base leading-relaxed">
          Dedikasi nyata santri Al-Hamid dalam menguasai turats, sains madrasah, diplomasi bahasa asing, serta seni budaya Islam.
        </p>

        <div class="space-y-4 pt-2">
          {#if data?.achievements}
            {#each data.achievements as a}
              <div class="p-5 rounded-2xl bg-slate-900/80 border border-slate-800 hover:border-amber-500/40 transition flex items-start gap-4">
                <div class="w-11 h-11 rounded-xl bg-amber-500/20 text-amber-400 flex items-center justify-center shrink-0">
                  <Trophy class="w-5 h-5" />
                </div>
                <div>
                  <div class="flex items-center gap-2">
                    <span class="text-[11px] font-black text-amber-400 uppercase tracking-wider">{a.level}</span>
                    <span class="text-xs text-slate-500">&bull; {a.year}</span>
                  </div>
                  <div class="font-display font-bold text-sm sm:text-base text-white mt-1 leading-snug">{a.title}</div>
                  <div class="text-xs text-slate-400 mt-1">{a.recipient_name}</div>
                </div>
              </div>
            {/each}
          {/if}
        </div>

        <div class="pt-2">
          <a
            href="/tentang/prestasi"
            class="inline-flex items-center gap-2 text-sm font-display font-bold text-amber-400 hover:text-amber-300 group"
          >
            <span>Lihat Direktori Prestasi Lengkap</span>
            <ArrowRight class="w-4 h-4 group-hover:translate-x-1.5 transition" />
          </a>
        </div>
      </div>

      <!-- Facilities Grid -->
      <div class="space-y-6">
        <div class="inline-flex items-center gap-2 px-3.5 py-1.5 rounded-full bg-teal-500/15 text-teal-300 text-xs font-bold border border-teal-500/30">
          <Building2 class="w-3.5 h-3.5 text-teal-400" />
          <span>Infrastruktur & Sarana</span>
        </div>

        <h2 class="font-display font-black text-3xl sm:text-5xl text-white tracking-tight">
          Sarana Representatif untuk Kenyamanan Santri
        </h2>

        <p class="text-slate-300 text-sm sm:text-base leading-relaxed">
          Ruang belajar modern, laboratorium komputer & sains, masjid megah, serta asrama higienis yang kondusif.
        </p>

        <div class="grid grid-cols-2 gap-4 pt-2">
          {#if data?.facilities}
            {#each data.facilities.slice(0, 4) as f}
              <div class="rounded-2xl overflow-hidden bg-slate-900 border border-slate-800 relative group aspect-[4/3]">
                <img
                  src={f.image_url}
                  alt={f.name}
                  class="w-full h-full object-cover group-hover:scale-105 transition duration-500"
                />
                <div class="absolute inset-0 bg-gradient-to-t from-slate-950 via-slate-950/40 to-transparent"></div>
                <div class="absolute bottom-3 left-3 right-3 text-white">
                  <span class="text-[10px] text-teal-400 uppercase tracking-widest font-black">{f.category}</span>
                  <div class="font-display font-bold text-xs sm:text-sm leading-snug line-clamp-1">{f.name}</div>
                </div>
              </div>
            {/each}
          {/if}
        </div>

        <div class="pt-2">
          <a
            href="/tentang/fasilitas"
            class="inline-flex items-center gap-2 text-sm font-display font-bold text-teal-400 hover:text-teal-300 group"
          >
            <span>Jelajahi Seluruh Fasilitas Kampus</span>
            <ArrowRight class="w-4 h-4 group-hover:translate-x-1.5 transition" />
          </a>
        </div>
      </div>
    </div>
  </div>
</section>

<!-- ================= 7. TESTIMONI & REPUTASI ================= -->
<section class="py-24 max-w-7xl mx-auto px-4 sm:px-8">
  <div class="text-center max-w-2xl mx-auto space-y-3 mb-16">
    <div class="inline-flex items-center gap-2 px-3.5 py-1.5 rounded-full bg-teal-50 border border-teal-200/70 text-teal-800 text-xs font-bold">
      Suara Nyata
    </div>
    <h2 class="font-display font-black text-3xl sm:text-5xl text-slate-950 tracking-tight">
      Apa Kata Alumni & Wali Santri
    </h2>
  </div>

  <div class="grid grid-cols-1 md:grid-cols-3 gap-8">
    {#if data?.testimonials}
      {#each data.testimonials as t}
        <div class="bg-white rounded-3xl border border-slate-200/80 p-8 shadow-sm hover:shadow-xl transition-all duration-300 flex flex-col justify-between space-y-6">
          <div class="space-y-4">
            <div class="flex items-center gap-1 text-amber-400">
              {#each Array(t.rating || 5) as _}
                <Star class="w-4 h-4 fill-amber-400" />
              {/each}
            </div>
            <p class="font-editorial italic text-slate-700 text-base leading-relaxed">
              "{t.content}"
            </p>
          </div>

          <div class="flex items-center gap-3.5 pt-4 border-t border-slate-100">
            <img
              src={t.avatar_url}
              alt={t.name}
              class="w-12 h-12 rounded-full object-cover border-2 border-teal-500 shadow-sm"
            />
            <div>
              <div class="font-display font-bold text-sm text-slate-950">{t.name}</div>
              <div class="text-xs text-slate-500">{t.role}</div>
            </div>
          </div>
        </div>
      {/each}
    {/if}
  </div>
</section>

<!-- ================= 8. ARCHITECTURAL CALL TO ACTION ================= -->
<section class="py-20 bg-gradient-to-r from-[#090D16] via-slate-950 to-[#0B1220] text-white relative overflow-hidden border-t border-slate-800/80">
  <div class="absolute -right-20 -bottom-20 w-[500px] h-[500px] bg-teal-500/15 rounded-full blur-3xl pointer-events-none"></div>

  <div class="max-w-7xl mx-auto px-4 sm:px-8 relative z-10 flex flex-col md:flex-row items-center justify-between gap-8">
    <div class="space-y-3 text-center md:text-left">
      <div class="inline-flex items-center gap-2 px-3.5 py-1 rounded-full bg-teal-500/15 text-teal-300 text-xs font-bold border border-teal-500/30">
        Penerimaan Santri Baru (PSB) 2026/2027
      </div>
      <h2 class="font-display font-black text-3xl sm:text-5xl text-white tracking-tight">
        Siapkan Masa Depan Putra-Putri Anda
      </h2>
      <p class="text-slate-300 text-sm sm:text-base max-w-xl">
        Kunjungi kampus kami di Cintamulya atau hubungi tim sekretariat kami untuk konsultasi jenjang pendidikan dan panduan pendaftaran.
      </p>
    </div>

    <div class="flex flex-wrap items-center gap-4 shrink-0">
      <a
        href="/kontak"
        class="px-8 py-4 rounded-full bg-gradient-to-r from-teal-400 to-emerald-400 hover:from-teal-300 hover:to-emerald-300 text-slate-950 font-display font-black text-sm tracking-wide shadow-xl shadow-teal-950/40 hover:scale-105 active:scale-95 transition duration-200"
      >
        Daftar Online Sekarang
      </a>
      <a
        href="https://wa.me/6281234567890"
        target="_blank"
        rel="noopener noreferrer"
        class="px-8 py-4 rounded-full bg-white/10 hover:bg-white/15 text-white font-display font-bold text-sm tracking-wide backdrop-blur-md border border-white/20 transition duration-200"
      >
        Konsultasi WhatsApp
      </a>
    </div>
  </div>
</section>
