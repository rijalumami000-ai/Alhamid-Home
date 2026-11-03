<script lang="ts">
  import { normalizeGoogleDriveUrl, isGoogleDriveUrl, extractGoogleDriveId } from '$lib/drive';
  import { Image, CheckCircle2, AlertCircle, ExternalLink } from 'lucide-svelte';

  interface Props {
    value: string;
    label?: string;
    id?: string;
    placeholder?: string;
    required?: boolean;
    helperText?: string;
  }

  let {
    value = $bindable(''),
    label = 'URL Gambar / Foto',
    id = 'image_url',
    placeholder = 'https://drive.google.com/... atau https://...',
    required = false,
    helperText = 'Mendukung link share Google Drive langsung (otomatis dikonversi ke gambar) atau link URL biasa.'
  }: Props = $props();

  let imgError = $state(false);
  let isDrive = $derived(isGoogleDriveUrl(value) || value.includes('googleusercontent.com/d/'));

  function handleInput(e: Event) {
    const input = e.target as HTMLInputElement;
    const raw = input.value;
    const normalized = normalizeGoogleDriveUrl(raw);
    value = normalized;
    imgError = false;
  }

  function handleBlur() {
    if (value) {
      value = normalizeGoogleDriveUrl(value);
    }
  }
</script>

<div class="space-y-2">
  <div class="flex items-center justify-between">
    <label for={id} class="text-xs font-bold text-slate-700 uppercase tracking-wider flex items-center gap-1.5">
      <Image class="w-3.5 h-3.5 text-teal-600" />
      <span>{label}</span>
      {#if required}<span class="text-rose-500 font-bold">*</span>{/if}
    </label>
    {#if isDrive}
      <span class="inline-flex items-center gap-1 text-[11px] font-semibold text-teal-700 bg-teal-50 px-2 py-0.5 rounded-full border border-teal-200">
        <CheckCircle2 class="w-3 h-3 text-teal-600" />
        Google Drive Terintegrasi
      </span>
    {/if}
  </div>

  <div class="relative">
    <input
      {id}
      type="url"
      {required}
      value={value}
      oninput={handleInput}
      onblur={handleBlur}
      {placeholder}
      class="w-full px-4 py-2.5 rounded-xl border border-slate-200 focus:border-teal-500 focus:ring-2 focus:ring-teal-500/20 text-sm font-mono text-slate-800 transition placeholder:text-slate-400"
    />
  </div>

  <p class="text-[11px] text-slate-500 leading-relaxed font-normal">
    {helperText}
  </p>

  <!-- Live Image Preview Card -->
  {#if value && value.trim().length > 0}
    <div class="mt-2.5 p-3 rounded-2xl bg-slate-50 border border-slate-200/90 flex flex-col sm:flex-row items-start sm:items-center gap-3">
      <div class="w-20 h-20 rounded-xl bg-slate-200 overflow-hidden shrink-0 border border-slate-300 relative group">
        {#if !imgError}
          <img
            src={value}
            alt="Preview URL"
            class="w-full h-full object-cover"
            onerror={() => (imgError = true)}
            onload={() => (imgError = false)}
          />
        {:else}
          <div class="w-full h-full flex flex-col items-center justify-center p-1 text-center bg-rose-50 text-rose-500 text-[10px]">
            <AlertCircle class="w-5 h-5 mb-1" />
            <span>Gagal muat</span>
          </div>
        {/if}
      </div>

      <div class="flex-1 min-w-0 space-y-1">
        <div class="flex items-center gap-2">
          <span class="text-xs font-bold text-slate-800">Preview Output Foto:</span>
          {#if !imgError}
            <span class="text-[11px] text-emerald-600 font-semibold flex items-center gap-1">
              <CheckCircle2 class="w-3 h-3" /> Siap Ditampilkan
            </span>
          {:else}
            <span class="text-[11px] text-rose-600 font-semibold flex items-center gap-1">
              <AlertCircle class="w-3 h-3" /> Gambar Tidak Terbuka
            </span>
          {/if}
        </div>

        <p class="text-[11px] text-slate-500 truncate font-mono">
          {value}
        </p>

        {#if imgError && isDrive}
          <p class="text-[11px] text-amber-700 bg-amber-50 p-1.5 rounded-lg border border-amber-200 font-medium">
            💡 <strong>Penting:</strong> Pastikan setelan berbagi di Google Drive telah diubah menjadi: <em>"Siapa saja yang memiliki link" (Anyone with the link can view)</em> agar gambar dapat diakses oleh publik.
          </p>
        {/if}
      </div>

      <a
        href={value}
        target="_blank"
        rel="noopener noreferrer"
        class="shrink-0 p-2 text-slate-500 hover:text-slate-900 hover:bg-slate-200/60 rounded-lg transition"
        title="Buka gambar di tab baru"
      >
        <ExternalLink class="w-4 h-4" />
      </a>
    </div>
  {/if}
</div>

