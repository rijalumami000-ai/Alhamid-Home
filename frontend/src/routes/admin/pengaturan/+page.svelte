<script lang="ts">
  import { onMount } from 'svelte';
  import { api } from '$lib/api';
  import { Settings, Save, CheckCircle2, Globe, Phone, Share2, BookOpen, Type, Check, Sparkles } from 'lucide-svelte';
  import { FONT_PRESETS, applyFontPreset, getSavedFontPresetId } from '$lib/fonts';
  import { updateLocalSettings } from '$lib/stores/settings';

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
  <title>Pengaturan Website | CMS Al-Hamid</title>
</svelte:head>

<div class="space-y-6 max-w-4xl">
  <div>
    <h1 class="font-serif-luxury font-bold text-2xl text-slate-900">
      Pengaturan Website Pesantren
    </h1>
    <p class="text-xs text-slate-500 mt-1">
      Atur identitas lembaga, kontak resmi, tautan media sosial, serta narasi profil dan visi-misi.
    </p>
  </div>

  {#if savedSuccess}
    <div class="p-4 rounded-2xl bg-emerald-50 border border-emerald-200 text-emerald-800 text-xs font-bold flex items-center gap-2">
      <CheckCircle2 class="w-4 h-4 text-emerald-600" />
      <span>Pengaturan website berhasil diperbarui dan langsung diterapkan ke halaman publik!</span>
    </div>
  {/if}

  <form onsubmit={handleSave} class="space-y-8">
    <!-- SECTION 1: IDENTITAS -->
    <div class="bg-white rounded-3xl p-6 sm:p-8 border border-slate-200 shadow-sm space-y-4">
      <div class="text-xs font-bold uppercase tracking-wider text-emerald-800 flex items-center gap-1.5 pb-2 border-b border-slate-100">
        <Globe class="w-4 h-4" /> Identitas Lembaga & Portal
      </div>

      <div class="space-y-1">
        <label for="s_name" class="text-xs font-bold text-slate-700 uppercase">Nama Pesantren *</label>
        <input
          id="s_name"
          type="text"
          bind:value={settings['site_name']}
          class="w-full px-4 py-2.5 rounded-xl border border-slate-200 text-sm font-medium"
        />
      </div>

      <div class="space-y-1">
        <label for="s_tagline" class="text-xs font-bold text-slate-700 uppercase">Tagline / Slogan</label>
        <input
          id="s_tagline"
          type="text"
          bind:value={settings['site_tagline']}
          class="w-full px-4 py-2.5 rounded-xl border border-slate-200 text-sm font-medium"
        />
      </div>

      <div class="space-y-1">
        <label for="s_desc" class="text-xs font-bold text-slate-700 uppercase">Deskripsi Ringkas (Meta SEO)</label>
        <textarea
          id="s_desc"
          rows="2"
          bind:value={settings['site_description']}
          class="w-full px-4 py-2.5 rounded-xl border border-slate-200 text-sm font-medium"
        ></textarea>
      </div>
    </div>

    <!-- SECTION 2: TIPOGRAFI & GAYA FONT WEBSITE -->
    <div class="bg-white rounded-3xl p-6 sm:p-8 border border-slate-200 shadow-sm space-y-6">
      <div class="flex flex-col sm:flex-row sm:items-center justify-between gap-2 pb-3 border-b border-slate-100">
        <div class="text-xs font-bold uppercase tracking-wider text-emerald-800 flex items-center gap-1.5">
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
            class="text-left p-5 rounded-2xl border-2 transition-all cursor-pointer relative flex flex-col justify-between space-y-4 {isSelected ? 'border-emerald-700 bg-emerald-50/40 shadow-md ring-2 ring-emerald-600/20' : 'border-slate-200/80 bg-slate-50/50 hover:border-slate-300 hover:bg-white'}"
          >
            <div class="space-y-2">
              <div class="flex items-center justify-between">
                <span class="text-[10px] font-mono uppercase tracking-widest font-semibold px-2 py-0.5 rounded-md {isSelected ? 'bg-emerald-700 text-white' : 'bg-slate-200 text-slate-700'}">
                  {preset.category}
                </span>
                {#if isSelected}
                  <span class="inline-flex items-center gap-1 text-[11px] font-bold text-emerald-800">
                    <Check class="w-3.5 h-3.5" /> Aktif
                  </span>
                {/if}
              </div>

              <div>
                <h3 class="font-bold text-base text-slate-900" style="font-family: {preset.fontDisplayCss};">
                  {preset.name}
                </h3>
                <p class="text-xs text-slate-500 mt-0.5">
                  {preset.tagline}
                </p>
              </div>

              <!-- Specimen Details -->
              <div class="pt-2 border-t border-slate-200/60 text-[11px] text-slate-600 space-y-1 font-mono">
                <div><span class="text-slate-400">Headings:</span> <strong class="text-slate-800">{preset.headingFont}</strong></div>
                <div><span class="text-slate-400">Body:</span> <strong class="text-slate-800">{preset.bodyFont}</strong></div>
                <div><span class="text-slate-400">Editorial:</span> <strong class="text-slate-800">{preset.editorialFont}</strong></div>
              </div>
            </div>

            <!-- Live Sample Preview -->
            <div class="p-3.5 rounded-xl bg-white border border-slate-200/70 shadow-2xs space-y-1.5">
              <div class="text-sm font-black text-slate-900 leading-snug" style="font-family: {preset.fontDisplayCss};">
                Pondok Pesantren Al-Hamid
              </div>
              <div class="text-xs text-slate-600 line-clamp-2 leading-relaxed" style="font-family: {preset.fontBodyCss};">
                {preset.previewQuote}
              </div>
              <div class="text-[11px] italic text-slate-500" style="font-family: {preset.fontEditorialCss};">
                "Kearifan turats, kemajuan sains peradaban."
              </div>
            </div>
          </div>
        {/each}
      </div>
    </div>

    <!-- SECTION 3: KONTAK & ALAMAT -->
    <div class="bg-white rounded-3xl p-6 sm:p-8 border border-slate-200 shadow-sm space-y-4">
      <div class="text-xs font-bold uppercase tracking-wider text-emerald-800 flex items-center gap-1.5 pb-2 border-b border-slate-100">
        <Phone class="w-4 h-4" /> Kontak & Lokasi Kampus
      </div>

      <div class="grid grid-cols-1 sm:grid-cols-2 gap-4">
        <div class="space-y-1">
          <label for="s_phone" class="text-xs font-bold text-slate-700 uppercase">Nomor Telepon</label>
          <input
            id="s_phone"
            type="text"
            bind:value={settings['site_phone']}
            class="w-full px-4 py-2.5 rounded-xl border border-slate-200 text-sm font-medium"
          />
        </div>
        <div class="space-y-1">
          <label for="s_wa" class="text-xs font-bold text-slate-700 uppercase">Nomor WhatsApp (Angka saja: 628...)</label>
          <input
            id="s_wa"
            type="text"
            bind:value={settings['site_whatsapp']}
            class="w-full px-4 py-2.5 rounded-xl border border-slate-200 text-sm font-medium"
          />
        </div>
      </div>

      <div class="space-y-1">
        <label for="s_email" class="text-xs font-bold text-slate-700 uppercase">Email Resmi</label>
        <input
          id="s_email"
          type="email"
          bind:value={settings['site_email']}
          class="w-full px-4 py-2.5 rounded-xl border border-slate-200 text-sm font-medium"
        />
      </div>

      <div class="space-y-1">
        <label for="s_addr" class="text-xs font-bold text-slate-700 uppercase">Alamat Lengkap</label>
        <textarea
          id="s_addr"
          rows="2"
          bind:value={settings['site_address']}
          class="w-full px-4 py-2.5 rounded-xl border border-slate-200 text-sm font-medium"
        ></textarea>
      </div>
    </div>

    <!-- SECTION 3: MEDIA SOSIAL -->
    <div class="bg-white rounded-3xl p-6 sm:p-8 border border-slate-200 shadow-sm space-y-4">
      <div class="text-xs font-bold uppercase tracking-wider text-emerald-800 flex items-center gap-1.5 pb-2 border-b border-slate-100">
        <Share2 class="w-4 h-4" /> Tautan Media Sosial
      </div>

      <div class="grid grid-cols-1 sm:grid-cols-2 gap-4">
        <div class="space-y-1">
          <label for="s_ig" class="text-xs font-bold text-slate-700 uppercase">URL Instagram</label>
          <input id="s_ig" type="url" bind:value={settings['social_instagram']} class="w-full px-4 py-2.5 rounded-xl border border-slate-200 text-sm font-medium" />
        </div>
        <div class="space-y-1">
          <label for="s_yt" class="text-xs font-bold text-slate-700 uppercase">URL YouTube</label>
          <input id="s_yt" type="url" bind:value={settings['social_youtube']} class="w-full px-4 py-2.5 rounded-xl border border-slate-200 text-sm font-medium" />
        </div>
        <div class="space-y-1">
          <label for="s_fb" class="text-xs font-bold text-slate-700 uppercase">URL Facebook</label>
          <input id="s_fb" type="url" bind:value={settings['social_facebook']} class="w-full px-4 py-2.5 rounded-xl border border-slate-200 text-sm font-medium" />
        </div>
        <div class="space-y-1">
          <label for="s_tt" class="text-xs font-bold text-slate-700 uppercase">URL TikTok</label>
          <input id="s_tt" type="url" bind:value={settings['social_tiktok']} class="w-full px-4 py-2.5 rounded-xl border border-slate-200 text-sm font-medium" />
        </div>
      </div>
    </div>

    <!-- SECTION 4: PROFIL, VISI, MISI -->
    <div class="bg-white rounded-3xl p-6 sm:p-8 border border-slate-200 shadow-sm space-y-4">
      <div class="text-xs font-bold uppercase tracking-wider text-emerald-800 flex items-center gap-1.5 pb-2 border-b border-slate-100">
        <BookOpen class="w-4 h-4" /> Narasi Profil, Visi & Misi
      </div>

      <div class="space-y-1">
        <label for="s_profile" class="text-xs font-bold text-slate-700 uppercase">Profil & Sejarah Pesantren</label>
        <textarea
          id="s_profile"
          rows="5"
          bind:value={settings['about_profile']}
          class="w-full px-4 py-2.5 rounded-xl border border-slate-200 text-sm font-medium leading-relaxed"
        ></textarea>
      </div>

      <div class="space-y-1">
        <label for="s_vision" class="text-xs font-bold text-slate-700 uppercase">Visi Pesantren</label>
        <textarea
          id="s_vision"
          rows="2"
          bind:value={settings['about_vision']}
          class="w-full px-4 py-2.5 rounded-xl border border-slate-200 text-sm font-medium leading-relaxed"
        ></textarea>
      </div>

      <div class="space-y-1">
        <label for="s_mission" class="text-xs font-bold text-slate-700 uppercase">Misi Pesantren (Daftar Poin)</label>
        <textarea
          id="s_mission"
          rows="5"
          bind:value={settings['about_mission']}
          class="w-full px-4 py-2.5 rounded-xl border border-slate-200 text-sm font-medium leading-relaxed"
        ></textarea>
      </div>
    </div>

    <!-- Submit Button -->
    <div class="flex justify-end">
      <button
        type="submit"
        class="px-8 py-3 rounded-xl bg-emerald-800 hover:bg-emerald-700 text-white font-bold text-sm shadow-lg transition flex items-center gap-2"
      >
        <Save class="w-4 h-4" />
        <span>Simpan Seluruh Pengaturan</span>
      </button>
    </div>
  </form>
</div>

