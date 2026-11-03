<script lang="ts">
  import { onMount } from 'svelte';
  import { api } from '$lib/api';
  import { Settings, Save, CheckCircle2, Globe, Type, Check, Sparkles, Layout } from 'lucide-svelte';
  import { FONT_PRESETS, applyFontPreset, getSavedFontPresetId } from '$lib/fonts';
  import { updateLocalSettings } from '$lib/stores/settings';
  import DriveImageInput from '$lib/components/DriveImageInput.svelte';

  let settings = $state<Record<string, string>>({});
  let loading = $state(true);
  let savedSuccess = $state(false);

  function selectFont(presetId: string) {
    settings['site_font_preset'] = presetId;
    applyFontPreset(presetId);
  }

  async function loadData() {
    try {
      const list = await api.admin.getSettings();
      const map: Record<string, string> = {};
      for (const item of list) {
        map[item.key] = item.value;
      }
      if (!map['site_font_preset']) {
        map['site_font_preset'] = getSavedFontPresetId() || 'architectural';
      }
      if (!map['header_brand_title']) map['header_brand_title'] = 'AL-HAMID';
      if (!map['header_brand_subtitle']) map['header_brand_subtitle'] = 'Pesantren & Academy • Cintamulya';
      if (!map['header_brand_initials']) map['header_brand_initials'] = 'AH';
      settings = map;
      applyFontPreset(settings['site_font_preset']);
    } catch (e) {
      console.error(e);
    } finally {
      loading = false;
    }
  }

  async function handleSave(e: SubmitEvent) {
    e.preventDefault();
    savedSuccess = false;
    try {
      await api.admin.updateSettings(settings);
      updateLocalSettings(settings);
      if (settings['site_font_preset']) {
        applyFontPreset(settings['site_font_preset']);
      }
      savedSuccess = true;
      setTimeout(() => (savedSuccess = false), 3500);
    } catch (err: any) {
      alert(err.message || 'Gagal menyimpan pengaturan');
    }
  }

  onMount(() => {
    loadData();
  });
</script>

<svelte:head>
  <title>Pengaturan Umum & Tampilan | CMS Al-Hamid</title>
</svelte:head>

<div class="space-y-6 max-w-4xl">
  <div>
    <h1 class="font-display font-black text-2xl text-slate-900 tracking-tight">
      Pengaturan Umum & Tampilan Website
    </h1>
    <p class="text-xs text-slate-500 mt-1">
      Atur identitas logo dan brand header web utama, nama lembaga, serta pilihan font dan tipografi global.
    </p>
  </div>

  {#if savedSuccess}
    <div class="p-4 rounded-2xl bg-emerald-50 border border-emerald-200 text-emerald-800 text-xs font-bold flex items-center gap-2 animate-in fade-in duration-200">
      <CheckCircle2 class="w-4 h-4 text-emerald-600" />
      <span>Pengaturan berhasil diperbarui dan langsung diterapkan ke seluruh website!</span>
    </div>
  {/if}

  <form onsubmit={handleSave} class="space-y-6">
    <!-- SECTION 1: HEADER BRAND IDENTITY (REQUEST #6) -->
    <div class="bg-white rounded-3xl p-6 sm:p-8 border border-slate-200 shadow-sm space-y-5">
      <div class="text-xs font-bold uppercase tracking-wider text-teal-800 flex items-center gap-1.5 pb-2 border-b border-slate-100">
        <Layout class="w-4 h-4" /> Identitas Brand Header (Pojok Header Web Utama)
      </div>

      <div class="grid grid-cols-1 sm:grid-cols-2 gap-4">
        <div class="space-y-1">
          <label for="h_title" class="text-xs font-bold text-slate-700 uppercase">Judul Brand Header *</label>
          <input
            id="h_title"
            type="text"
            required
            bind:value={settings['header_brand_title']}
            placeholder="Contoh: AL-HAMID"
            class="w-full px-4 py-2.5 rounded-xl border border-slate-200 text-sm font-bold uppercase tracking-tight"
          />
        </div>

        <div class="space-y-1">
          <label for="h_sub" class="text-xs font-bold text-slate-700 uppercase">Subjudul / Tagline Header *</label>
          <input
            id="h_sub"
            type="text"
            required
            bind:value={settings['header_brand_subtitle']}
            placeholder="Contoh: Pesantren & Academy • Cintamulya"
            class="w-full px-4 py-2.5 rounded-xl border border-slate-200 text-sm font-medium"
          />
        </div>
      </div>

      <div class="grid grid-cols-1 sm:grid-cols-3 gap-4 items-start">
        <div class="space-y-1 sm:col-span-1">
          <label for="h_init" class="text-xs font-bold text-slate-700 uppercase">Inisial Logo (Fallback)</label>
          <input
            id="h_init"
            type="text"
            maxlength="4"
            bind:value={settings['header_brand_initials']}
            placeholder="AH"
            class="w-full px-4 py-2.5 rounded-xl border border-slate-200 text-sm font-black text-center font-display"
          />
          <p class="text-[11px] text-slate-400">Tampil jika logo gambar tidak diisi.</p>
        </div>

        <div class="sm:col-span-2">
          <DriveImageInput
            id="h_logo"
            label="URL Foto / Logo Icon Header"
            bind:value={settings['header_brand_logo_url']}
            placeholder="https://drive.google.com/... atau URL gambar logo"
            helperText="Mendukung link share Google Drive langsung. Jika diisi, kotak inisial digantikan logo ini."
          />
        </div>
      </div>

      <!-- Live Header Brand Preview -->
      <div class="mt-3 p-4 rounded-2xl bg-slate-50 border border-slate-200/80 flex items-center justify-between">
        <div class="flex items-center gap-3">
          <div class="w-11 h-11 rounded-xl bg-slate-950 border border-slate-800 text-teal-400 flex items-center justify-center font-display font-black text-xl shadow-md overflow-hidden">
            {#if settings['header_brand_logo_url'] && settings['header_brand_logo_url'].trim().length > 0}
              <img src={settings['header_brand_logo_url']} alt="Preview Logo" class="w-full h-full object-cover" />
            {:else}
              {settings['header_brand_initials'] || 'AH'}
            {/if}
          </div>
          <div>
            <div class="font-display font-extrabold text-lg text-slate-950 tracking-tight leading-none">
              {settings['header_brand_title'] || 'AL-HAMID'}
            </div>
            <div class="text-[10px] font-bold tracking-widest text-slate-400 uppercase mt-1">
              {settings['header_brand_subtitle'] || 'Pesantren & Academy • Cintamulya'}
            </div>
          </div>
        </div>
        <span class="text-[10px] font-mono text-slate-400 bg-white px-2.5 py-1 rounded-md border border-slate-200">
          Live Preview Header
        </span>
      </div>
    </div>

    <!-- SECTION 2: IDENTITAS PORTAL & SEO -->
    <div class="bg-white rounded-3xl p-6 sm:p-8 border border-slate-200 shadow-sm space-y-4">
      <div class="text-xs font-bold uppercase tracking-wider text-teal-800 flex items-center gap-1.5 pb-2 border-b border-slate-100">
        <Globe class="w-4 h-4" /> Identitas Lembaga & Meta SEO
      </div>

      <div class="space-y-1">
        <label for="s_name" class="text-xs font-bold text-slate-700 uppercase">Nama Lembaga Pesantren *</label>
        <input
          id="s_name"
          type="text"
          required
          bind:value={settings['site_name']}
          class="w-full px-4 py-2.5 rounded-xl border border-slate-200 text-sm font-medium"
        />
      </div>

      <div class="space-y-1">
        <label for="s_tagline" class="text-xs font-bold text-slate-700 uppercase">Tagline / Slogan Resmi</label>
        <input
          id="s_tagline"
          type="text"
          bind:value={settings['site_tagline']}
          class="w-full px-4 py-2.5 rounded-xl border border-slate-200 text-sm font-medium"
        />
      </div>

      <div class="space-y-1">
        <label for="s_desc" class="text-xs font-bold text-slate-700 uppercase">Deskripsi Ringkas (Meta Description)</label>
        <textarea
          id="s_desc"
          rows="2"
          bind:value={settings['site_description']}
          class="w-full px-4 py-2.5 rounded-xl border border-slate-200 text-sm font-medium leading-relaxed"
        ></textarea>
      </div>
    </div>

    <!-- SECTION 3: TIPOGRAFI & GAYA FONT WEBSITE -->
    <div class="bg-white rounded-3xl p-6 sm:p-8 border border-slate-200 shadow-sm space-y-6">
      <div class="flex flex-col sm:flex-row sm:items-center justify-between gap-2 pb-3 border-b border-slate-100">
        <div class="text-xs font-bold uppercase tracking-wider text-teal-800 flex items-center gap-1.5">
          <Type class="w-4 h-4" /> Tipografi & Presensi Font Global
        </div>
        <span class="text-[11px] text-slate-500">
          Pilihan langsung mengubah gaya tipografi di seluruh halaman publik dan CMS
        </span>
      </div>

      <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
        {#each FONT_PRESETS as preset}
          {@const isSelected = (settings['site_font_preset'] || 'architectural') === preset.id}
          <div
            role="button"
            tabindex="0"
            onclick={() => selectFont(preset.id)}
            onkeydown={(e) => e.key === 'Enter' && selectFont(preset.id)}
            class="text-left p-5 rounded-2xl border-2 transition-all cursor-pointer relative flex flex-col justify-between space-y-4 {isSelected ? 'border-teal-700 bg-teal-50/40 shadow-md ring-2 ring-teal-600/20' : 'border-slate-200/80 bg-slate-50/50 hover:border-slate-300 hover:bg-white'}"
          >
            <div class="space-y-2">
              <div class="flex items-center justify-between">
                <span class="text-[10px] font-mono uppercase tracking-widest font-semibold px-2 py-0.5 rounded-md {isSelected ? 'bg-teal-700 text-white' : 'bg-slate-200 text-slate-700'}">
                  {preset.category}
                </span>
                {#if isSelected}
                  <span class="inline-flex items-center gap-1 text-[11px] font-bold text-teal-800">
                    <Check class="w-3.5 h-3.5" /> Aktif
                  </span>
                {/if}
              </div>

              <div>
                <h3 class="font-bold text-base text-slate-900" style="font-family: {preset.fontDisplayCss};">
                  {preset.name}
                </h3>
                <p class="text-xs text-slate-500 mt-1">{preset.description}</p>
              </div>

              <div class="p-3 rounded-xl bg-white/80 border border-slate-200/60 space-y-1.5">
                <div class="text-sm font-bold text-slate-900 truncate" style="font-family: {preset.fontDisplayCss};">
                  Pesantren Modern Berkarakter & Unggul
                </div>
                <div class="text-xs text-slate-600 line-clamp-1" style="font-family: {preset.fontBodyCss};">
                  Menjaga tradisi sanad kitab kuning dan keunggulan sains masa kini.
                </div>
              </div>
            </div>

            <div class="text-[11px] text-slate-400 font-mono">
              Font: <span class="text-slate-700 font-bold">{preset.displayFont}</span> &bull; {preset.bodyFont}
            </div>
          </div>
        {/each}
      </div>
    </div>

    <!-- Submit Button -->
    <div class="flex justify-end pt-2">
      <button
        type="submit"
        class="inline-flex items-center gap-2 px-8 py-3 rounded-xl bg-teal-700 hover:bg-teal-800 text-white font-bold text-xs shadow-lg shadow-teal-900/20 transition active:scale-95"
      >
        <Save class="w-4 h-4" />
        <span>Simpan Pengaturan</span>
      </button>
    </div>
  </form>
</div>
