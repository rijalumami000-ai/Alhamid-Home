<script lang="ts">
  import { onMount } from 'svelte';
  import { api } from '$lib/api';
  import type { MediaPhoto, MediaVideo, Category } from '$lib/types';
  import { Plus, Edit2, Trash2, X, Play, Image, Video, ArrowUpDown } from 'lucide-svelte';
  import ConfirmDeleteModal from '$lib/components/ConfirmDeleteModal.svelte';
  import DriveImageInput from '$lib/components/DriveImageInput.svelte';

  let activeTab = $state<'photos' | 'videos'>('photos');
  let photos = $state<MediaPhoto[]>([]);
  let videos = $state<MediaVideo[]>([]);
  let photoCategories = $state<Category[]>([]);
  let videoCategories = $state<Category[]>([]);

  // Modals & Editing states
  let isPhotoModalOpen = $state(false);
  let editingPhoto = $state<MediaPhoto | null>(null);
  let photoForm = $state({
    title: '',
    category_id: 1,
    image_url: '',
    caption: '',
    order: 0,
  });

  let isVideoModalOpen = $state(false);
  let editingVideo = $state<MediaVideo | null>(null);
  let videoForm = $state({
    title: '',
    category_id: 1,
    youtube_url: '',
    description: '',
    order: 0,
  });

  // Delete modal state
  let isDeleteModalOpen = $state(false);
  let deleteTarget = $state<{ type: 'photo' | 'video'; id: number; name: string } | null>(null);
  let isDeleting = $state(false);

  async function loadData() {
    try {
      const [pList, vList, pCats, vCats] = await Promise.all([
        api.admin.getPhotos(),
        api.admin.getVideos(),
        api.getCategories('galeri_foto'),
        api.getCategories('galeri_video'),
      ]);
      photos = pList;
      videos = vList;
      photoCategories = pCats;
      videoCategories = vCats;
      if (pCats.length > 0 && !editingPhoto) photoForm.category_id = pCats[0].id;
      if (vCats.length > 0 && !editingVideo) videoForm.category_id = vCats[0].id;
    } catch (e) {
      console.error(e);
    }
  }

  // --- Photo Actions ---
  function openCreatePhoto() {
    editingPhoto = null;
    photoForm = {
      title: '',
      category_id: photoCategories[0]?.id || 1,
      image_url: 'https://images.unsplash.com/photo-1542838132-92c53300491e?auto=format&fit=crop&w=1000&q=80',
      caption: '',
      order: photos.length + 1,
    };
    isPhotoModalOpen = true;
  }

  function openEditPhoto(p: MediaPhoto) {
    editingPhoto = p;
    photoForm = {
      title: p.title,
      category_id: p.category_id,
      image_url: p.image_url,
      caption: p.caption,
      order: p.order || 0,
    };
    isPhotoModalOpen = true;
  }

  async function handleSavePhoto(e: SubmitEvent) {
    e.preventDefault();
    try {
      if (editingPhoto) {
        await api.admin.updatePhoto(editingPhoto.id, photoForm);
      } else {
        await api.admin.createPhoto(photoForm);
      }
      isPhotoModalOpen = false;
      await loadData();
    } catch (err: any) {
      alert(err.message || 'Gagal menyimpan foto');
    }
  }

  function promptDeletePhoto(p: MediaPhoto) {
    deleteTarget = { type: 'photo', id: p.id, name: p.title };
    isDeleteModalOpen = true;
  }

  // --- Video Actions ---
  function openCreateVideo() {
    editingVideo = null;
    videoForm = {
      title: '',
      category_id: videoCategories[0]?.id || 1,
      youtube_url: 'https://www.youtube.com/watch?v=dQw4w9WgXcQ',
      description: '',
      order: videos.length + 1,
    };
    isVideoModalOpen = true;
  }

  function openEditVideo(v: MediaVideo) {
    editingVideo = v;
    videoForm = {
      title: v.title,
      category_id: v.category_id,
      youtube_url: v.youtube_url,
      description: v.description,
      order: v.order || 0,
    };
    isVideoModalOpen = true;
  }

  async function handleSaveVideo(e: SubmitEvent) {
    e.preventDefault();
    try {
      if (editingVideo) {
        await api.admin.updateVideo(editingVideo.id, videoForm);
      } else {
        await api.admin.createVideo(videoForm);
      }
      isVideoModalOpen = false;
      await loadData();
    } catch (err: any) {
      alert(err.message || 'Gagal menyimpan video');
    }
  }

  function promptDeleteVideo(v: MediaVideo) {
    deleteTarget = { type: 'video', id: v.id, name: v.title };
    isDeleteModalOpen = true;
  }

  async function handleConfirmDelete() {
    if (!deleteTarget) return;
    isDeleting = true;
    try {
      if (deleteTarget.type === 'photo') {
        await api.admin.deletePhoto(deleteTarget.id);
      } else {
        await api.admin.deleteVideo(deleteTarget.id);
      }
      isDeleteModalOpen = false;
      deleteTarget = null;
      await loadData();
    } catch (err: any) {
      alert(err.message || 'Gagal menghapus');
    } finally {
      isDeleting = false;
    }
  }

  onMount(() => {
    loadData();
  });
</script>

<svelte:head>
  <title>Kelola Galeri Media | CMS Al-Hamid</title>
</svelte:head>

<div class="space-y-6">
  <div class="flex flex-col sm:flex-row sm:items-center justify-between gap-4">
    <div>
      <h1 class="font-display font-black text-2xl text-slate-900 tracking-tight">
        Galeri Foto & Video YouTube
      </h1>
      <p class="text-xs text-slate-500 mt-1">
        Kelola album dokumentasi foto kegiatan santri dan sematan video pengajian resmi dengan pengaturan urutan tampil.
      </p>
    </div>
    <div>
      {#if activeTab === 'photos'}
        <button
          onclick={openCreatePhoto}
          class="inline-flex items-center gap-1.5 px-4 py-2.5 rounded-xl bg-teal-700 hover:bg-teal-800 text-white text-xs font-bold shadow-md transition active:scale-95"
        >
          <Plus class="w-4 h-4" /> Tambah Foto
        </button>
      {:else}
        <button
          onclick={openCreateVideo}
          class="inline-flex items-center gap-1.5 px-4 py-2.5 rounded-xl bg-teal-700 hover:bg-teal-800 text-white text-xs font-bold shadow-md transition active:scale-95"
        >
          <Plus class="w-4 h-4" /> Tambah Video YouTube
        </button>
      {/if}
    </div>
  </div>

  <!-- Tabs -->
  <div class="flex items-center gap-2 border-b border-slate-200">
    <button
      onclick={() => (activeTab = 'photos')}
      class="px-5 py-3 text-xs font-bold border-b-2 transition {activeTab === 'photos'
        ? 'border-teal-700 text-teal-800'
        : 'border-transparent text-slate-500 hover:text-slate-800'}"
    >
      Galeri Foto ({photos.length})
    </button>
    <button
      onclick={() => (activeTab = 'videos')}
      class="px-5 py-3 text-xs font-bold border-b-2 transition {activeTab === 'videos'
        ? 'border-teal-700 text-teal-800'
        : 'border-transparent text-slate-500 hover:text-slate-800'}"
    >
      Galeri Video ({videos.length})
    </button>
  </div>

  <!-- PHOTOS TAB -->
  {#if activeTab === 'photos'}
    <div class="grid grid-cols-2 sm:grid-cols-3 md:grid-cols-4 gap-4">
      {#each photos as p}
        <div class="bg-white rounded-2xl border border-slate-200 overflow-hidden shadow-sm flex flex-col justify-between group">
          <div class="relative aspect-square bg-slate-900">
            <img src={p.image_url} alt={p.title} class="w-full h-full object-cover" />
            <div class="absolute top-2 left-2 flex items-center gap-1.5">
              <span class="px-2 py-0.5 rounded bg-slate-950/80 text-white text-[9px] font-bold uppercase">
                {p.category?.name}
              </span>
              <span class="px-2 py-0.5 rounded bg-teal-800/90 text-white text-[9px] font-mono font-bold">
                #{p.order || 0}
              </span>
            </div>
          </div>
          <div class="p-3 space-y-1">
            <div class="font-bold text-xs text-slate-900 line-clamp-1">{p.title}</div>
            <div class="flex items-center justify-between pt-2 border-t border-slate-100">
              <span class="text-[10px] text-slate-400">
                {new Date(p.created_at).toLocaleDateString('id-ID')}
              </span>
              <div class="flex items-center gap-1">
                <button
                  onclick={() => openEditPhoto(p)}
                  class="text-slate-500 hover:text-teal-700 hover:bg-slate-100 p-1.5 rounded-lg transition"
                  title="Edit Foto"
                >
                  <Edit2 class="w-3.5 h-3.5" />
                </button>
                <button
                  onclick={() => promptDeletePhoto(p)}
                  class="text-rose-500 hover:text-rose-700 hover:bg-rose-50 p-1.5 rounded-lg transition"
                  title="Hapus Foto"
                >
                  <Trash2 class="w-3.5 h-3.5" />
                </button>
              </div>
            </div>
          </div>
        </div>
      {/each}
    </div>
  {/if}

  <!-- VIDEOS TAB -->
  {#if activeTab === 'videos'}
    <div class="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 gap-6">
      {#each videos as v}
        <div class="bg-white rounded-3xl border border-slate-200 overflow-hidden shadow-sm flex flex-col justify-between">
          <div>
            <div class="aspect-video bg-slate-900 relative">
              <iframe
                src="https://www.youtube.com/embed/{v.youtube_id}"
                title={v.title}
                class="w-full h-full"
                allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture"
                allowfullscreen
              ></iframe>
            </div>
            <div class="p-4 space-y-1">
              <div class="flex items-center gap-2">
                <span class="px-2 py-0.5 rounded bg-teal-50 text-teal-800 border border-teal-200 text-[10px] font-bold uppercase">
                  {v.category?.name}
                </span>
                <span class="px-2 py-0.5 rounded bg-slate-100 text-slate-700 text-[10px] font-mono font-bold">
                  Urutan #{v.order || 0}
                </span>
              </div>
              <h3 class="font-display font-bold text-sm text-slate-900 line-clamp-2 pt-1">{v.title}</h3>
              <p class="text-xs text-slate-500 line-clamp-2">{v.description}</p>
            </div>
          </div>
          <div class="p-4 pt-0 flex justify-end gap-1">
            <button
              onclick={() => openEditVideo(v)}
              class="p-2 rounded-xl text-slate-500 hover:text-teal-700 hover:bg-slate-100 transition"
              title="Edit Video"
            >
              <Edit2 class="w-4 h-4" />
            </button>
            <button
              onclick={() => promptDeleteVideo(v)}
              class="p-2 rounded-xl text-slate-500 hover:text-rose-600 hover:bg-rose-50 transition"
              title="Hapus Video"
            >
              <Trash2 class="w-4 h-4" />
            </button>
          </div>
        </div>
      {/each}
    </div>
  {/if}
</div>

<!-- PHOTO MODAL -->
{#if isPhotoModalOpen}
  <div class="fixed inset-0 z-50 bg-slate-950/60 backdrop-blur-sm flex items-center justify-center p-4">
    <div class="bg-white w-full max-w-md rounded-3xl shadow-2xl border border-slate-200 overflow-hidden">
      <div class="p-6 border-b border-slate-100 flex items-center justify-between">
        <h2 class="font-display font-bold text-lg text-slate-900">
          {editingPhoto ? 'Edit Foto Galeri' : 'Tambah Foto ke Galeri'}
        </h2>
        <button onclick={() => (isPhotoModalOpen = false)}><X class="w-5 h-5 text-slate-400" /></button>
      </div>
      <form onsubmit={handleSavePhoto} class="p-6 space-y-4 max-h-[80vh] overflow-y-auto">
        <div class="space-y-1">
          <label for="p_title" class="text-xs font-bold text-slate-700 uppercase">Judul Foto *</label>
          <input
            id="p_title"
            type="text"
            required
            bind:value={photoForm.title}
            placeholder="Momen pengajian..."
            class="w-full px-4 py-2.5 rounded-xl border border-slate-200 text-sm font-medium"
          />
        </div>

        <div class="grid grid-cols-2 gap-3">
          <div class="space-y-1">
            <label for="p_cat" class="text-xs font-bold text-slate-700 uppercase">Kategori Album</label>
            <select
              id="p_cat"
              bind:value={photoForm.category_id}
              class="w-full px-4 py-2.5 rounded-xl border border-slate-200 text-sm font-medium bg-white"
            >
              {#each photoCategories as c}
                <option value={c.id}>{c.name}</option>
              {/each}
            </select>
          </div>
          <div class="space-y-1">
            <label for="p_order" class="text-xs font-bold text-slate-700 uppercase">Urutan Tampil</label>
            <input
              id="p_order"
              type="number"
              bind:value={photoForm.order}
              class="w-full px-4 py-2.5 rounded-xl border border-slate-200 text-sm font-medium"
            />
          </div>
        </div>

        <DriveImageInput
          id="p_img"
          label="File Foto Dokumentasi"
          required
          bind:value={photoForm.image_url}
          placeholder="https://drive.google.com/... atau https://..."
        />

        <div class="space-y-1">
          <label for="p_cap" class="text-xs font-bold text-slate-700 uppercase">Keterangan / Kepsyen</label>
          <textarea
            id="p_cap"
            rows="2"
            bind:value={photoForm.caption}
            class="w-full px-4 py-2.5 rounded-xl border border-slate-200 text-sm font-medium"
          ></textarea>
        </div>
        <div class="pt-4 border-t border-slate-100 flex justify-end gap-2">
          <button type="button" onclick={() => (isPhotoModalOpen = false)} class="px-4 py-2 rounded-xl border border-slate-200 text-xs font-bold">Batal</button>
          <button type="submit" class="px-5 py-2 rounded-xl bg-teal-700 hover:bg-teal-800 text-white text-xs font-bold shadow-md transition">
            {editingPhoto ? 'Perbarui Foto' : 'Simpan Foto'}
          </button>
        </div>
      </form>
    </div>
  </div>
{/if}

<!-- VIDEO MODAL -->
{#if isVideoModalOpen}
  <div class="fixed inset-0 z-50 bg-slate-950/60 backdrop-blur-sm flex items-center justify-center p-4">
    <div class="bg-white w-full max-w-md rounded-3xl shadow-2xl border border-slate-200 overflow-hidden">
      <div class="p-6 border-b border-slate-100 flex items-center justify-between">
        <h2 class="font-display font-bold text-lg text-slate-900">
          {editingVideo ? 'Edit Video YouTube' : 'Tambah Video YouTube'}
        </h2>
        <button onclick={() => (isVideoModalOpen = false)}><X class="w-5 h-5 text-slate-400" /></button>
      </div>
      <form onsubmit={handleSaveVideo} class="p-6 space-y-4 max-h-[80vh] overflow-y-auto">
        <div class="space-y-1">
          <label for="v_title" class="text-xs font-bold text-slate-700 uppercase">Judul Video *</label>
          <input
            id="v_title"
            type="text"
            required
            bind:value={videoForm.title}
            placeholder="Kajian Kitab..."
            class="w-full px-4 py-2.5 rounded-xl border border-slate-200 text-sm font-medium"
          />
        </div>
        <div class="grid grid-cols-2 gap-3">
          <div class="space-y-1">
            <label for="v_cat" class="text-xs font-bold text-slate-700 uppercase">Kategori Video</label>
            <select
              id="v_cat"
              bind:value={videoForm.category_id}
              class="w-full px-4 py-2.5 rounded-xl border border-slate-200 text-sm font-medium bg-white"
            >
              {#each videoCategories as c}
                <option value={c.id}>{c.name}</option>
              {/each}
            </select>
          </div>
          <div class="space-y-1">
            <label for="v_order" class="text-xs font-bold text-slate-700 uppercase">Urutan Tampil</label>
            <input
              id="v_order"
              type="number"
              bind:value={videoForm.order}
              class="w-full px-4 py-2.5 rounded-xl border border-slate-200 text-sm font-medium"
            />
          </div>
        </div>
        <div class="space-y-1">
          <label for="v_url" class="text-xs font-bold text-slate-700 uppercase">Tautan / URL YouTube *</label>
          <input
            id="v_url"
            type="url"
            required
            bind:value={videoForm.youtube_url}
            placeholder="https://www.youtube.com/watch?v=..."
            class="w-full px-4 py-2.5 rounded-xl border border-slate-200 text-sm font-medium"
          />
        </div>
        <div class="space-y-1">
          <label for="v_desc" class="text-xs font-bold text-slate-700 uppercase">Deskripsi Singkat</label>
          <textarea
            id="v_desc"
            rows="2"
            bind:value={videoForm.description}
            class="w-full px-4 py-2.5 rounded-xl border border-slate-200 text-sm font-medium"
          ></textarea>
        </div>
        <div class="pt-4 border-t border-slate-100 flex justify-end gap-2">
          <button type="button" onclick={() => (isVideoModalOpen = false)} class="px-4 py-2 rounded-xl border border-slate-200 text-xs font-bold">Batal</button>
          <button type="submit" class="px-5 py-2 rounded-xl bg-teal-700 hover:bg-teal-800 text-white text-xs font-bold shadow-md transition">
            {editingVideo ? 'Perbarui Video' : 'Simpan Video'}
          </button>
        </div>
      </form>
    </div>
  </div>
{/if}

<!-- CONFIRM DELETE MODAL -->
<ConfirmDeleteModal
  isOpen={isDeleteModalOpen}
  title={deleteTarget?.type === 'photo' ? 'Hapus Foto Galeri' : 'Hapus Video'}
  itemName={deleteTarget?.name}
  message={deleteTarget?.type === 'photo'
    ? 'Foto ini akan dihapus dari album galeri dokumentasi pesantren.'
    : 'Video ini akan dihapus dari daftar putar galeri pesantren.'}
  {isDeleting}
  onConfirm={handleConfirmDelete}
  onCancel={() => {
    isDeleteModalOpen = false;
    deleteTarget = null;
  }}
/>
