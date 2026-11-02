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
    ArrowUpRight
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
    }, 6000);

    return () => clearInterval(interval);
  });
</script>

<svelte:head>
  <title>Pondok Pesantren Al-Hamid | Berakhlak, Unggul & Berdaya Saing Global</title>
  <meta name="description" content="Pondok Pesantren Al-Hamid Cintamulya: Paduan harmoni antara tradisi luhur kitab kuning dan keunggulan sains modern (MTs, MA, MAI, Diniyah, PMQ)." />
</svelte:head>

<!-- ================= 1. HERO SECTION WITH BACKGROUND SLIDER ================= -->
<section class="relative min-h-[640px] lg:min-h-[720px] flex items-center bg-slate-950 text-white overflow-hidden">
  {#if data && data.banners.length > 0}
    {#each data.banners as banner, i}
      <div
        class="absolute inset-0 transition-opacity duration-1000 ease-in-out {i === activeBannerIndex ? 'opacity-100 z-10' : 'opacity-0 z-0'}"
      >
        <!-- Background Image with Overlay -->
        <img
          src={banner.image_url}
          alt={banner.title}
          class="w-full h-full object-cover object-center scale-105 transform transition-transform duration-10000"
        />
        <div class="absolute inset-0 bg-gradient-to-r from-slate-950/90 via-slate-950/70 to-transparent"></div>
        <div class="absolute inset-0 bg-emerald-950/30 mix-blend-multiply"></div>
      </div>
    {/each}

    <!-- Content overlay -->
    <div class="relative z-20 max-w-7xl mx-auto px-4 sm:px-8 py-20 w-full">
      <div class="max-w-2xl space-y-6">
        <!-- Badge -->
        {#if data.banners[activeBannerIndex]?.badge}
          <div class="inline-flex items-center gap-2 px-3.5 py-1.5 rounded-full bg-emerald-500/20 border border-emerald-400/40 text-emerald-300 text-xs font-semibold backdrop-blur-md">
            <Sparkles class="w-3.5 h-3.5 text-amber-400" />
            <span>{data.banners[activeBannerIndex].badge}</span>
          </div>
        {/if}

        <!-- Title -->
        <h1 class="font-serif-luxury font-bold text-3xl sm:text-5xl lg:text-6xl text-white tracking-tight leading-[1.15]">
          {data.banners[activeBannerIndex]?.title}
        </h1>

        <!-- Subtitle -->
        <p class="text-base sm:text-lg text-slate-200 leading-relaxed max-w-xl">
          {data.banners[activeBannerIndex]?.subtitle}
        </p>

        <!-- CTA Buttons -->
        <div class="pt-4 flex flex-wrap items-center gap-4">
          <a
            href={data.banners[activeBannerIndex]?.button_link || '/kontak'}
            class="px-7 py-3.5 rounded-xl bg-gradient-to-r from-emerald-600 to-emerald-500 hover:from-emerald-500 hover:to-emerald-400 text-white font-semibold text-sm shadow-xl shadow-emerald-950/40 hover:scale-105 transition duration-200 flex items-center gap-2"
          >
            <span>{data.banners[activeBannerIndex]?.button_text || 'Jelajahi Sekarang'}</span>
            <ArrowRight class="w-4 h-4" />
          </a>
          <a
            href="/pendidikan"
            class="px-7 py-3.5 rounded-xl bg-white/10 hover:bg-white/20 text-white font-semibold text-sm backdrop-blur-md border border-white/20 transition duration-200 flex items-center gap-2"
          >
            <span>Program Pendidikan</span>
            <Compass class="w-4 h-4" />
          </a>
        </div>
      </div>

      <!-- Slide Indicators -->
      <div class="absolute bottom-8 left-4 sm:left-8 flex items-center gap-2 z-20">
        {#each data.banners as _, i}
          <button
            onclick={() => (activeBannerIndex = i)}
            class="h-2 rounded-full transition-all duration-300 {i === activeBannerIndex ? 'w-8 bg-amber-400' : 'w-2 bg-white/40'}"
            aria-label="Slide {i + 1}"
          ></button>
        {/each}
      </div>
    </div>
  {:else}
    <div class="relative z-20 max-w-7xl mx-auto px-4 sm:px-8 py-24 text-center">
      <div class="inline-flex items-center gap-2 px-3.5 py-1.5 rounded-full bg-emerald-500/20 text-emerald-300 text-xs font-semibold mb-4">
        Pondok Pesantren Al-Hamid
      </div>
      <h1 class="font-serif-luxury font-bold text-4xl sm:text-6xl text-white mb-6">
        Menempa Akhlak, Mengukir Prestasi Masa Depan
      </h1>
      <p class="text-slate-300 max-w-xl mx-auto mb-8">
        Paduan harmoni antara kedalaman turats pesantren salaf dan keunggulan sains modern.
      </p>
    </div>
  {/if}
</section>

<!-- ================= 2. QUICK STATS BAR ================= -->
<section class="relative z-30 -mt-10 max-w-7xl mx-auto px-4 sm:px-8">
  <div class="bg-white rounded-2xl shadow-xl border border-slate-100 p-6 sm:p-8 grid grid-cols-2 lg:grid-cols-4 gap-6 text-center">
    <div class="space-y-1">
      <div class="font-serif-luxury font-bold text-3xl sm:text-4xl text-emerald-800">
        {data?.stats?.santri_count || '1.450+'}
      </div>
      <div class="text-xs sm:text-sm font-medium text-slate-500">Santri Mukim & Aktif</div>
    </div>
    <div class="space-y-1 border-l border-slate-100 pl-4 sm:pl-6">
      <div class="font-serif-luxury font-bold text-3xl sm:text-4xl text-emerald-800">
        {data?.stats?.asatidz_count || '85+'}
      </div>
      <div class="text-xs sm:text-sm font-medium text-slate-500">Dewan Asatidz & Masyayikh</div>
    </div>
    <div class="space-y-1 border-l-0 lg:border-l border-slate-100 pl-0 lg:pl-6">
      <div class="font-serif-luxury font-bold text-3xl sm:text-4xl text-emerald-800">
        {data?.stats?.alumni_count || '6.200+'}
      </div>
      <div class="text-xs sm:text-sm font-medium text-slate-500">Alumni Tersebar Nasional</div>
    </div>
    <div class="space-y-1 border-l border-slate-100 pl-4 sm:pl-6">
      <div class="font-serif-luxury font-bold text-3xl sm:text-4xl text-emerald-800">
        {data?.stats?.prestasi_count || '120+'}
      </div>
      <div class="text-xs sm:text-sm font-medium text-slate-500">Penghargaan & Prestasi</div>
    </div>
  </div>
</section>

<!-- ================= 3. SAMBUTAN PENGASUH & PROFIL RINGKAS ================= -->
<section class="py-20 lg:py-28 max-w-7xl mx-auto px-4 sm:px-8">
  <div class="grid grid-cols-1 lg:grid-cols-12 gap-12 items-center">
    <!-- Image Card -->
    <div class="lg:col-span-5 relative">
      <div class="relative rounded-3xl overflow-hidden shadow-2xl border-4 border-white aspect-[4/5]">
        <img
          src="https://images.unsplash.com/photo-1542838132-92c53300491e?auto=format&fit=crop&w=1000&q=80"
          alt="Suasana Pengajian Kitab Pondok Pesantren Al-Hamid"
          class="w-full h-full object-cover"
        />
        <div class="absolute inset-0 bg-gradient-to-t from-slate-950/80 via-transparent to-transparent"></div>
        <div class="absolute bottom-6 left-6 right-6 text-white">
          <div class="text-xs uppercase tracking-wider text-amber-300 font-semibold mb-1">
            Mursyid & Pengasuh Utama
          </div>
          <div class="font-serif-luxury font-bold text-xl">K.H. Ahmad Hamidurrohman</div>
          <p class="text-xs text-slate-300 mt-1">
            "Menjaga kemurnian sanad, merawat tradisi, menatap masa depan peradaban."
          </p>
        </div>
      </div>
      <!-- Floating Accent Badge -->
      <div class="absolute -bottom-4 -right-4 bg-emerald-800 text-white p-4 rounded-2xl shadow-xl flex items-center gap-3 hidden sm:flex">
        <CheckCircle2 class="w-8 h-8 text-amber-300 shrink-0" />
        <div class="text-xs">
          <div class="font-bold">Sanad Muttashil</div>
          <div class="text-emerald-200">Masyayikh Lirboyo & Sarang</div>
        </div>
      </div>
    </div>

    <!-- Text Description -->
    <div class="lg:col-span-7 space-y-6">
      <div class="inline-flex items-center gap-2 px-3.5 py-1.5 rounded-full bg-emerald-100 text-emerald-800 text-xs font-semibold">
        <Sparkles class="w-3.5 h-3.5 text-emerald-700" /> Profil Pondok Pesantren Al-Hamid
      </div>
      <h2 class="font-serif-luxury font-bold text-3xl sm:text-4xl lg:text-5xl text-slate-900 leading-tight">
        Melahirkan Generasi Unggul Berakhlak Mulia & Berdaya Saing Global
      </h2>
      <p class="text-slate-600 text-base leading-relaxed">
        Pondok Pesantren Al-Hamid didirikan sebagai wadah kaderisasi ulama dan intelektual muslim. Kami meyakini bahwa kedalaman spiritualitas dan pemahaman teks-teks klasik (turats) harus berjalan seiring dengan kecakapan sains, teknologi informasi, serta komunikasi global.
      </p>
      <div class="grid grid-cols-1 sm:grid-cols-2 gap-4 pt-2">
        <div class="p-4 rounded-2xl bg-white border border-slate-100 shadow-sm flex items-start gap-3">
          <BookOpen class="w-5 h-5 text-emerald-600 shrink-0 mt-0.5" />
          <div>
            <div class="font-semibold text-sm text-slate-800">Turats & Tahfidz</div>
            <div class="text-xs text-slate-500 mt-0.5">Madrasah Diniyah Salafiyah & PMQ bersanad talaqqi.</div>
          </div>
        </div>
        <div class="p-4 rounded-2xl bg-white border border-slate-100 shadow-sm flex items-start gap-3">
          <GraduationCap class="w-5 h-5 text-emerald-600 shrink-0 mt-0.5" />
          <div>
            <div class="font-semibold text-sm text-slate-800">Pendidikan Formal</div>
            <div class="text-xs text-slate-500 mt-0.5">MTs, MA Cintamulya (Akreditasi A), & MAI Kejuruan IT.</div>
          </div>
        </div>
      </div>
      <div class="pt-2 flex items-center gap-4">
        <a
          href="/tentang/profil"
          class="inline-flex items-center gap-2 text-sm font-semibold text-emerald-800 hover:text-emerald-700 group"
        >
          <span>Baca Selengkapnya Profil & Sejarah</span>
          <ArrowRight class="w-4 h-4 group-hover:translate-x-1 transition" />
        </a>
      </div>
    </div>
  </div>
</section>

<!-- ================= 4. PENDIDIKAN UNGGULAN (FORMAL & NON-FORMAL) ================= -->
<section class="py-20 bg-slate-100/70 border-y border-slate-200/60">
  <div class="max-w-7xl mx-auto px-4 sm:px-8">
    <div class="text-center max-w-2xl mx-auto space-y-3 mb-14">
      <div class="inline-flex items-center gap-2 px-3.5 py-1.5 rounded-full bg-emerald-100 text-emerald-800 text-xs font-semibold">
        Jenjang Pendidikan
      </div>
      <h2 class="font-serif-luxury font-bold text-3xl sm:text-4xl text-slate-900">
        Ekosistem Pendidikan Komprehensif
      </h2>
      <p class="text-slate-600 text-sm sm:text-base">
        Dari pondasi gramatika kitab kuning hingga teknologi rekayasa informatika, seluruh jenjang dirancang terpadu untuk kemandirian santri.
      </p>
    </div>

    <!-- Grid of Units -->
    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
      {#if data?.education_units}
        {#each data.education_units as unit}
          <div class="bg-white rounded-2xl border border-slate-100 shadow-md hover:shadow-xl transition-all duration-300 overflow-hidden flex flex-col group">
            <div class="relative h-48 overflow-hidden bg-slate-900">
              <img
                src={unit.hero_image}
                alt={unit.name}
                class="w-full h-full object-cover group-hover:scale-105 transition duration-500"
              />
              <div class="absolute inset-0 bg-gradient-to-t from-slate-950/80 via-transparent to-transparent"></div>
              <div class="absolute top-3 right-3">
                <span class="px-2.5 py-1 rounded-full text-[11px] font-bold tracking-wide uppercase shadow-sm {unit.type === 'formal' ? 'bg-blue-600 text-white' : 'bg-emerald-700 text-white'}">
                  {unit.type === 'formal' ? 'Pendidikan Formal' : 'Pendidikan Non-Formal'}
                </span>
              </div>
              <div class="absolute bottom-3 left-4 right-4">
                <span class="text-amber-300 text-xs font-medium">{unit.badge}</span>
              </div>
            </div>

            <div class="p-6 flex-1 flex flex-col justify-between space-y-4">
              <div>
                <h3 class="font-serif-luxury font-bold text-xl text-slate-900 group-hover:text-emerald-700 transition">
                  {unit.name}
                </h3>
                <p class="text-slate-500 text-xs mt-1 font-medium">{unit.tagline}</p>
                <p class="text-slate-600 text-sm mt-3 line-clamp-3 leading-relaxed">
                  {unit.short_desc}
                </p>
              </div>

              <div class="pt-2 border-t border-slate-100 flex items-center justify-between">
                <span class="text-xs text-slate-400">Kepala: {unit.headmaster_name}</span>
                <a
                  href="/pendidikan/{unit.slug}"
                  class="inline-flex items-center gap-1.5 text-xs font-bold text-emerald-700 hover:text-emerald-800 transition"
                >
                  <span>Detail Laman</span>
                  <ArrowRight class="w-3.5 h-3.5" />
                </a>
              </div>
            </div>
          </div>
        {/each}
      {/if}
    </div>
  </div>
</section>

<!-- ================= 5. HIGHLIGHT BERITA & ARTIKEL ================= -->
<section class="py-20 lg:py-24 max-w-7xl mx-auto px-4 sm:px-8">
  <div class="flex flex-col md:flex-row md:items-end justify-between gap-4 mb-12">
    <div>
      <div class="inline-flex items-center gap-2 px-3.5 py-1.5 rounded-full bg-emerald-100 text-emerald-800 text-xs font-semibold mb-3">
        Kabar & Literasi
      </div>
      <h2 class="font-serif-luxury font-bold text-3xl sm:text-4xl text-slate-900">
        Kajian & Dinamika Pesantren Terkini
      </h2>
    </div>
    <a
      href="/berita"
      class="inline-flex items-center gap-2 text-sm font-semibold text-emerald-800 hover:text-emerald-700 group"
    >
      <span>Lihat Semua Berita</span>
      <ArrowRight class="w-4 h-4 group-hover:translate-x-1 transition" />
    </a>
  </div>

  <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-8">
    {#if data?.featured_articles}
      {#each data.featured_articles as article}
        <article class="bg-white rounded-2xl border border-slate-100 shadow-sm hover:shadow-lg transition-all duration-300 overflow-hidden flex flex-col group">
          <a href="/berita/{article.slug}" class="relative h-52 overflow-hidden bg-slate-100 block">
            <img
              src={article.thumbnail_url}
              alt={article.title}
              class="w-full h-full object-cover group-hover:scale-105 transition duration-500"
            />
            <div class="absolute top-3 left-3">
              <span class="px-2.5 py-1 rounded-full bg-emerald-800/90 text-white text-[11px] font-semibold backdrop-blur-sm">
                {article.category?.name || 'Kajian'}
              </span>
            </div>
          </a>

          <div class="p-6 flex-1 flex flex-col justify-between space-y-3">
            <div>
              <div class="text-xs text-slate-400 mb-2 flex items-center gap-2">
                <Calendar class="w-3.5 h-3.5" />
                <span>{new Date(article.published_at).toLocaleDateString('id-ID', { day: 'numeric', month: 'long', year: 'numeric' })}</span>
                <span>&bull;</span>
                <span>{article.author}</span>
              </div>
              <h3 class="font-serif-luxury font-bold text-lg text-slate-900 group-hover:text-emerald-700 transition line-clamp-2">
                <a href="/berita/{article.slug}">{article.title}</a>
              </h3>
              <p class="text-slate-600 text-sm mt-2 line-clamp-2 leading-relaxed">
                {article.excerpt}
              </p>
            </div>

            <div class="pt-3 border-t border-slate-100 flex items-center justify-between">
              <span class="text-xs text-slate-400">{article.view_count} pembaca</span>
              <a
                href="/berita/{article.slug}"
                class="inline-flex items-center gap-1 text-xs font-bold text-emerald-700 hover:text-emerald-800"
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

<!-- ================= 6. PRESTASI & FASILITAS SHOWCASE ================= -->
<section class="py-20 bg-slate-950 text-white">
  <div class="max-w-7xl mx-auto px-4 sm:px-8">
    <div class="grid grid-cols-1 lg:grid-cols-2 gap-12 items-center">
      <!-- Achievements List -->
      <div class="space-y-6">
        <div class="inline-flex items-center gap-2 px-3.5 py-1.5 rounded-full bg-amber-500/20 text-amber-300 text-xs font-semibold border border-amber-500/30">
          <Trophy class="w-3.5 h-3.5 text-amber-400" /> Prestasi Membanggakan
        </div>
        <h2 class="font-serif-luxury font-bold text-3xl sm:text-4xl text-white">
          Jejak Prestasi Santri di Panggung Nasional
        </h2>
        <p class="text-slate-300 text-sm sm:text-base leading-relaxed">
          Bukti nyata dedikasi santri Al-Hamid mengukir prestasi gemilang di bidang qira'atil kutub, sains madrasah, pidato multi bahasa, dan seni hadroh.
        </p>

        <div class="space-y-4 pt-2">
          {#if data?.achievements}
            {#each data.achievements as a}
              <div class="p-4 rounded-2xl bg-slate-900/80 border border-slate-800 flex items-start gap-4">
                <div class="w-10 h-10 rounded-xl bg-amber-500/20 text-amber-400 flex items-center justify-center shrink-0">
                  <Trophy class="w-5 h-5" />
                </div>
                <div>
                  <div class="flex items-center gap-2">
                    <span class="text-xs font-bold text-amber-400 uppercase tracking-wide">{a.level}</span>
                    <span class="text-xs text-slate-500">&bull; {a.year}</span>
                  </div>
                  <div class="font-semibold text-sm text-white mt-0.5">{a.title}</div>
                  <div class="text-xs text-slate-400 mt-1">{a.recipient_name}</div>
                </div>
              </div>
            {/each}
          {/if}
        </div>

        <div class="pt-2">
          <a
            href="/tentang/prestasi"
            class="inline-flex items-center gap-2 text-sm font-semibold text-amber-300 hover:text-amber-200"
          >
            <span>Lihat Galeri Prestasi Lengkap</span>
            <ArrowRight class="w-4 h-4" />
          </a>
        </div>
      </div>

      <!-- Facilities Preview Grid -->
      <div class="space-y-6">
        <div class="inline-flex items-center gap-2 px-3.5 py-1.5 rounded-full bg-emerald-500/20 text-emerald-300 text-xs font-semibold border border-emerald-500/30">
          <Building2 class="w-3.5 h-3.5 text-emerald-400" /> Sarana & Prasarana
        </div>
        <h2 class="font-serif-luxury font-bold text-3xl sm:text-4xl text-white">
          Fasilitas Representatif untuk Kenyamanan Belajar
        </h2>
        <div class="grid grid-cols-2 gap-4">
          {#if data?.facilities}
            {#each data.facilities.slice(0, 4) as f}
              <div class="rounded-2xl overflow-hidden bg-slate-900 border border-slate-800 relative group aspect-[4/3]">
                <img
                  src={f.image_url}
                  alt={f.name}
                  class="w-full h-full object-cover group-hover:scale-105 transition duration-500"
                />
                <div class="absolute inset-0 bg-gradient-to-t from-slate-950/90 via-slate-950/40 to-transparent"></div>
                <div class="absolute bottom-3 left-3 right-3 text-white">
                  <span class="text-[10px] text-emerald-400 uppercase tracking-wider font-bold">{f.category}</span>
                  <div class="font-semibold text-xs leading-snug line-clamp-1">{f.name}</div>
                </div>
              </div>
            {/each}
          {/if}
        </div>
        <div class="pt-2">
          <a
            href="/tentang/fasilitas"
            class="inline-flex items-center gap-2 text-sm font-semibold text-emerald-300 hover:text-emerald-200"
          >
            <span>Eksplorasi Seluruh Fasilitas</span>
            <ArrowRight class="w-4 h-4" />
          </a>
        </div>
      </div>
    </div>
  </div>
</section>

<!-- ================= 7. TESTIMONI SANTRI & WALI ================= -->
<section class="py-20 max-w-7xl mx-auto px-4 sm:px-8">
  <div class="text-center max-w-2xl mx-auto space-y-3 mb-14">
    <div class="inline-flex items-center gap-2 px-3.5 py-1.5 rounded-full bg-emerald-100 text-emerald-800 text-xs font-semibold">
      Suara Pengalaman
    </div>
    <h2 class="font-serif-luxury font-bold text-3xl sm:text-4xl text-slate-900">
      Apa Kata Alumni & Wali Santri
    </h2>
  </div>

  <div class="grid grid-cols-1 md:grid-cols-3 gap-6">
    {#if data?.testimonials}
      {#each data.testimonials as t}
        <div class="bg-white rounded-2xl border border-slate-100 p-6 shadow-sm hover:shadow-md transition flex flex-col justify-between space-y-4">
          <div class="space-y-3">
            <div class="flex items-center gap-1 text-amber-400">
              {#each Array(t.rating || 5) as _}
                <Star class="w-4 h-4 fill-amber-400" />
              {/each}
            </div>
            <p class="text-slate-600 text-sm leading-relaxed italic">
              "{t.content}"
            </p>
          </div>

          <div class="flex items-center gap-3 pt-3 border-t border-slate-100">
            <img
              src={t.avatar_url}
              alt={t.name}
              class="w-11 h-11 rounded-full object-cover border-2 border-emerald-600"
            />
            <div>
              <div class="font-bold text-sm text-slate-800">{t.name}</div>
              <div class="text-xs text-slate-500">{t.role}</div>
            </div>
          </div>
        </div>
      {/each}
    {/if}
  </div>
</section>

<!-- ================= 8. CALL TO ACTION PENDAFTARAN ================= -->
<section class="py-16 bg-gradient-to-tr from-emerald-950 via-emerald-900 to-emerald-800 text-white relative overflow-hidden">
  <div class="max-w-7xl mx-auto px-4 sm:px-8 relative z-10 flex flex-col md:flex-row items-center justify-between gap-8">
    <div class="space-y-3 text-center md:text-left">
      <span class="inline-flex items-center gap-1.5 px-3 py-1 rounded-full bg-emerald-800/80 text-amber-300 font-semibold text-xs border border-emerald-700">
        Penerimaan Santri Baru (PSB) 2026/2027
      </span>
      <h2 class="font-serif-luxury font-bold text-3xl sm:text-4xl text-white">
        Siapkan Masa Depan Putra-Putri Anda Bersama Al-Hamid
      </h2>
      <p class="text-emerald-100/80 text-sm sm:text-base max-w-xl">
        Kunjungi kampus kami di Cintamulya atau hubungi tim sekretariat kami untuk konsultasi jenjang pendidikan dan panduan pendaftaran.
      </p>
    </div>

    <div class="flex flex-wrap items-center gap-4 shrink-0">
      <a
        href="/kontak"
        class="px-8 py-3.5 rounded-xl bg-amber-400 hover:bg-amber-300 text-slate-950 font-bold text-sm shadow-xl transition duration-200"
      >
        Daftar Online Sekarang
      </a>
      <a
        href="https://wa.me/6281234567890"
        target="_blank"
        rel="noopener noreferrer"
        class="px-8 py-3.5 rounded-xl bg-white/10 hover:bg-white/20 text-white font-semibold text-sm backdrop-blur-md border border-white/20 transition duration-200"
      >
        Chat WhatsApp
      </a>
    </div>
  </div>
</section>
